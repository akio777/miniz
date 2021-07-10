package database

import (
	"backend/functions"
	"database/sql"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const (
	init_table = `
		CREATE TABLE IF NOT EXISTS users(
			id SERIAL NOT NULL,
			username varchar(30) NOT NULL,
			password varchar(255) NOT NULL,
			created_date TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_date TIMESTAMP NOT NULL DEFAULT NOW(),
			PRIMARY KEY (username)
		)
	`
	insert = `
		INSERT INTO users(username, password, created_date, updated_date)
		VALUES ($1, $2, NOW(), NOW());
	`
)

func Init_user(db *sql.DB) {
	_, err := db.Exec(init_table)
	if err != nil {
		fmt.Printf("Error from create users table :  %v\n", err)
		panic(err)
	}
	defer db.Close()
}

func Create_user(username string, password string) error {
	db := Call_DB()
	defer db.Close()
	pwd := functions.BytePassword(password)
	hashPwd, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return errors.New("password process error")
	}
	_, err = db.Exec(insert, username, string(hashPwd))
	if err != nil {
		return errors.New("username already exists")
	}
	return nil
}
