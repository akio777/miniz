package svc

import (
	"backend/models"
	"fmt"
	"github.com/go-pg/pg/v10"
	"golang.org/x/crypto/bcrypt"
	"log"
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

func (u *_User) Init(DB *pg.DB) error {
	_, err := DB.Exec(_init)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}

func (u *_User) Create(DB *pg.DB, user *models.Users) error {

	pwd := BytePassword(user.Password)
	fmt.Println(pwd)
	hashPwd, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	user.Password = string(hashPwd)
	if err != nil {
		return err
	}
	if _, err := DB.Model(user).Insert(); err != nil {
		return fmt.Errorf("username already exists")
	}
	return nil
}

func (u *_User) ReadByUsername(DB *pg.DB, username string) (*models.Users, error) {
	user := new(models.Users)
	if err := DB.Model(user).Where("username=?", username).Select(); err != nil {
		return nil, err
	}
	return user, nil

	// row := DB.QueryRow(fmt.Sprintf(_read, username))
}
