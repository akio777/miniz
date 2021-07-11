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
