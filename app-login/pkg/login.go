package pkg

import (
	"database/sql"
	"framework/api"
	"net/http"
)

func ProcessLogin(db *sql.DB, email string, password string) (*http.Cookie, error) {
	query := `
	SELECT password_hash
	  FROM bugbase_user
	 WHERE email = $1;
	`
	// Retrieve the user's password hash from the database
	var passwordHash string
	err := db.QueryRow(query, email).Scan(&passwordHash)
	if err != nil {
		return nil, err
	}

	// Create a new JwtHandler
	jwtHandler := api.NewJwtHandler("jwt", []byte("superDuperSecret"))

	// Process the login
	cookie, err := jwtHandler.ProcessLogin(email, password, passwordHash)
	if err != nil {
		return nil, err
	}

	return cookie, nil
}
