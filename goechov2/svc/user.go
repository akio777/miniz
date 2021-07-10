package svc

import (
	"backend/models"
	"database/sql"
	"fmt"
	// "golang.org/x/crypto/bcrypt"
)

type _User struct{}

const (
	_insert = `
		INSERT INTO user_info(username, password, created_date, updated_date)
		VALUES ($1, $2, NOW(), NOW());
	`
)

func (u *_User) Create(DB *sql.DB, user *models.User_info) error {

	// pwd := BytePassword(user.Password)

	// hashPwd, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	// if err != nil {
	// 	return err
	// }
	// if _, err = DB.Exec(_insert, user.Username, string(hashPwd)); err != nil {
	// 	return err
	// }
	sql_prepare, err := DB.Prepare(_insert)
	if err != nil {
		fmt.Println("Error : ", err)
		return err
	}
	fmt.Println(sql_prepare)
	defer DB.Close()
	return nil
}
