package models

import (
	"database/sql"
)

type User_info struct {
	Id       int64  `json:"id" xml:"id"`
	Username string `json:"username" xml:"username"`
	Password string `json:"password" xml:"password"`
	Group_id int64  `json:"group_id" xml:"group_id"`
	Role_id  int64  `json:"role_id" xml:"role_id"`
}

type Role_info struct {
	Id   int64  `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
}

type Group_info struct {
	Id           int64        `json:"id" xml:"id"`
	Name         string       `json:"name" xml:"name"`
	Created_date sql.NullTime `json:"created_date" xml:"created_date"`
	Updated_date sql.NullTime `json:"updated_date" xml:"updated_date"`
	Work_id      int64        `json:"work_id" xml:"work_id"`
}

type Work_info struct {
	Id           int64        `json:"id" xml:"id"`
	Name         string       `json:"name" xml:"name"`
	Created_date sql.NullTime `json:"created_date" xml:"created_date"`
	Updated_date sql.NullTime `json:"updated_date" xml:"updated_date"`
}
