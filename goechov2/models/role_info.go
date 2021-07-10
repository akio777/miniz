package models

import (
	"time"
)

type Role_info struct {
	Id           int64        `json:"id" xml:"id"`
	Name         string       `json:"name" xml:"name"`
	Created_date time.Time `json:"created_date" xml:"created_date"`
	Updated_date time.Time `json:"updated_date" xml:"updated_date"`
}
