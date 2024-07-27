package models

import (
	"api/db"
	"api/utils"
	"errors"
	"fmt"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := fmt.Sprintf("INSERT INTO %s (email, password) VALUES (?, ?)", db.TABLE_USERS)

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

	u.Id, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func (u *User) Validate() error {
	query := fmt.Sprintf(`SELECT id, password FROM %s WHERE email = ?`, db.TABLE_USERS)

	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.Id, &retrievedPassword)
	if err != nil {
		return err
	}

	validPassword := utils.CheckPassword(u.Password, retrievedPassword)

	if !validPassword {
		return errors.New("invalid credentials")
	}
	return nil
}
