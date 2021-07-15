package models

import "time"

type Users struct {
	tableName    struct{}  `sql:"users"`
	Id           int64     `json:"id" xml:"id" sql:"id"`
	Username     string    `json:"username" xml:"username" sql:"username, pk"`
	Password     string    `json:"password" xml:"password" sql:"password"`
	Created_date time.Time `json:"created_date" xml:"created_date" sql:"created_date"`
	Updated_date time.Time `json:"updated_date" xml:"updated_date" sql:"updated_date"`
}
