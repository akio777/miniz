package svc

import (
	"backend/models"
	"database/sql"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type _User struct{}

const (
	_init = `
		CREATE TABLE IF NOT EXISTS users(
			id SERIAL NOT NULL,
			username varchar(30) NOT NULL,
			password varchar(255) NOT NULL,
			created_date TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_date TIMESTAMP NOT NULL DEFAULT NOW(),
			PRIMARY KEY (username)
		)
	`

	_insert = `
		INSERT INTO users(username, password, created_date, updated_date)
		VALUES ($1, $2, NOW(), NOW());
	`

	_read = `
		SELECT * FROM users WHERE username = '%v'
	`
)

func (u *_User) Init(DB *sql.DB) error {
	_, err := DB.Exec(_init)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}

func (u *_User) Create(DB *sql.DB, user *models.User_info) error {

	pwd := BytePassword(user.Password)

	hashPwd, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	if _, err = DB.Exec(_insert, user.Username, string(hashPwd)); err != nil {
		return fmt.Errorf("username already exists")
	}
	return nil
}

func (u *_User) Read(DB *sql.DB, username string) ([]models.User_info, error) {
	users := []models.User_info{}
	rows, err := DB.Query(fmt.Sprintf(_read, username))
	if err != nil {
		return users, err
	}
	for rows.Next() {
	}

	return users, nil
}
