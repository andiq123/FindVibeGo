package userRepository

import (
	"database/sql"
	"github.com/google/uuid"
)

var DB *sql.DB

func GetOrCreateUser(name string) (userID string, err error) {
	query := `SELECT id FROM users WHERE name = ?`
	stmt, err := DB.Prepare(query)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	err = stmt.QueryRow(name).Scan(&userID)
	if err == nil {
		// User already exists, return the existing user ID
		return userID, nil
	}
	if err != sql.ErrNoRows {
		// Some other error occurred
		return "", err
	}

	// User doesn't exist, create a new one
	query = `INSERT INTO users (id, name) VALUES (?,?)`
	stmt, err = DB.Prepare(query)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	userID = uuid.New().String()
	_, err = stmt.Exec(userID, name)
	if err != nil {
		return "", err
	}

	return userID, nil
}
