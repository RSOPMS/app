package pkg

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
)

func ReadUserProfile(db *sql.DB, email string) (*User, error) {
	query := `
	SELECT bugbase_user.id,
		   bugbase_user.name,
		   bugbase_user.surname,
		   role.name,
		   bugbase_user.email
	  FROM bugbase_user
	  JOIN role ON bugbase_user.role_id = role.id
	 WHERE bugbase_user.email = $1;
	`

	user := &User{}
	err := db.QueryRow(query, email).Scan(
		&user.Id,
		&user.Name,
		&user.Surname,
		&user.Role,
		&user.Email,
	)

	if err != nil {
		return nil, err
	}

	// Generate unique avatar URL
	seed := fmt.Sprintf("%s%s%s", user.Name, user.Surname, user.Email)
	avatarURL := fmt.Sprintf("https://api.dicebear.com/9.x/pixel-art/svg?seed=%s", seed)

	// Fetch the avatar SVG image
	resp, err := http.Get(avatarURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the SVG image content
	avatarSVG, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	user.AvatarSVG = base64.StdEncoding.EncodeToString(avatarSVG)

	return user, nil
}
