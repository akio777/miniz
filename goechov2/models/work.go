package models

import (
	"database/sql"
)

type Work_info struct {
	Id           int64        `json:"id" xml:"id"`
	Name         string       `json:"name" xml:"name"`
	Created_date sql.NullTime `json:"created_date" xml:"created_date"`
	Updated_date sql.NullTime `json:"updated_date" xml:"updated_date"`
}
