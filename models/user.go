package models

import (
	"errors"
	"restapi/db"
	"restapi/utils"
)

type User struct {
	ID      int64
	Name    string
	Email   string `binding:"required"`
	Pasword string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?,?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(u.Pasword)

	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	u.ID = userId
	return err
}

func (u User) ValidateCredentials() error {
	query := "SELECT id, mail, password FROM users WHERE email =?"
	row := db.DB.QueryRow(query, u.Email)
	var retrievedPassword string

	err := row.Scan(&u.ID, &retrievedPassword)
	if err != nil {

	}
	passwordIsValid := utils.CheckPassword(u.Pasword, retrievedPassword)
	if !passwordIsValid {
		return errors.New("Credentials invalid")
	}
	return nil

}
