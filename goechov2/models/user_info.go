package models

import "time"

type User_info struct {
	Id           int64     `json:"id" xml:"id"`
	Username     string    `json:"username" xml:"username"`
	Password     string    `json:"password" xml:"password"`
	Group_id     int64     `json:"group_id" xml:"group_id"`
	Created_date time.Time `json:"created_date" xml:"created_date"`
	Updated_date time.Time `json:"updated_date" xml:"updated_date"`
}
