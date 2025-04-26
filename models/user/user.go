package user

import (
	"errors"

	"example.com/udemy_course/db"
	"example.com/udemy_course/utils"
)

type User struct {
	Id int64 `json:"id"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}


func Create(u User) error {
	query := `
	INSERT INTO User(Email, Password) VALUES
	(?, ?)
	`

	stmt, err := db.DB.Prepare(query)

	
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(u.Email, u.Password)

	return err
}

func Read(id int64) (*User, error) {
	query := `
	SELECT * FROM User WHERE Id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)

	var u User
	err = row.Scan(&u.Id, &u.Email, &u.Password)

	if err != nil {
		return nil, err
	}

	
	return &u, nil
}

func ReadAll() ([]User, error) {
	return nil, nil
}

func Update(u User) error {
	return nil
}

func Delete(id int) error {
	return nil
}

func ValidateUser(u User) (string, error) {
	query := `
	SELECT Id, Password From User Where Email = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return "", err
	}

	defer stmt.Close()

	row := stmt.QueryRow(u.Email)

	var  receivedPassword string
	err = row.Scan(&u.Id, &receivedPassword)

	if err != nil {
		return "", err
	}

	validated := utils.CheckHashPassword(u.Password, receivedPassword)

	if !validated {
		return "", errors.New("Invalid")
	} 
	
	return utils.GenerateToken(u.Email, u.Id)
}