package models

import (
	"errors"

	"example.com/app/db"
	"example.com/app/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(
		user.Email,
		hashedPassword,
	)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	user.ID = id
	return err
}

func (user User) ValidateCredentials() error {
	query := "SELECT password FROM users WHERE email = ?"

	row := db.DB.QueryRow(query, user.Email)

	var hashedPassword string

	err := row.Scan(&hashedPassword)

	if err != nil {
		return err
	}

	isPasswordValid := utils.VerifyPasswordHash(hashedPassword, user.Password)

	if !isPasswordValid {
		return errors.New("invalid credentials")
	}

	return nil
}
