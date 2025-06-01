package models

import (
	"errors"

	"example.com/event-booking-rest-api/db"
	"example.com/event-booking-rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
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

func GetUserById(id int64) (*User, error) {
	query := "SELECT * FROM users where id = ?"

	row := db.DB.QueryRow(query, id)

	var user User
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users where email = ?"

	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return errors.New("credentials Invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)
	if !passwordIsValid {
		return errors.New("credentials Invalid")
	}

	return nil
}
