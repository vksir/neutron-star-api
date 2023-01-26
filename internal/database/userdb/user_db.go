package userdb

import (
	"neutron-star-api/internal/database"
)

type User struct {
	ID                string `json:"id"`
	Username          string `json:"username"`
	EncryptedPassword string `json:"encrypted_password"`
}

func Get(username string) (*User, error) {
	u := User{}
	row := database.DB.QueryRow(
		"select * from t_user where username = $1", username)
	if err := row.Scan(&u.ID, &u.Username, &u.EncryptedPassword); err != nil {
		return nil, err
	}
	return &u, nil
}

func Set(username, encryptedPassword string) error {
	_, err := database.DB.Exec(
		"insert into t_user (username, encrypted_password) values ($1, $2)", username, encryptedPassword)
	return err
}
