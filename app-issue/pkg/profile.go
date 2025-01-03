package pkg

import (
	"database/sql"
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

	return user, nil
}
