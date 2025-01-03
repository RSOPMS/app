package pkg

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type Role struct {
	Id   int
	Name string
}

func ProcessRegister(db *sql.DB, name string, surname string, roleId int, email string, password string) error {
	query := `
	INSERT INTO bugbase_user (email, password_hash, name, surname, role_id, avatar_svg)
	  VALUES ($1, $2, $3, $4, $5, $6);
	`

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Generate unique avatar URL
	seed := fmt.Sprintf("%s%s%s", name, surname, email)
	avatarURL := fmt.Sprintf("https://api.dicebear.com/9.x/pixel-art/svg?seed=%s", seed)

	// Fetch the avatar SVG image
	resp, err := http.Get(avatarURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read the SVG image content
	avatarSVG, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	avatarSVGBase64 := base64.StdEncoding.EncodeToString(avatarSVG)

	_, err = db.Exec(query, email, string(passwordHash), name, surname, roleId, avatarSVGBase64)

	return err
}

func ReadRoles(db *sql.DB) ([]*Role, error) {
	query := `
	SELECT id, name
	  FROM role
	 ORDER BY id;
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	roles := []*Role{}
	for rows.Next() {
		role := &Role{}
		rows.Scan(&role.Id, &role.Name)
		roles = append(roles, role)
	}

	return roles, err
}
