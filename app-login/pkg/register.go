package pkg

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type Role struct {
	Id   int
	Name string
}

func ProcessRegister(db *sql.DB, name string, surname string, roleId int, email string, password string) error {
	query := `
	INSERT INTO bugbase_user (email, password_hash, name, surname, role_id)
	  VALUES ($1, $2, $3, $4, $5);
	`

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = db.Exec(query, email, string(passwordHash), name, surname, roleId)

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
