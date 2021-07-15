package models

import (
	"time"
)

type Group_info struct {
	Id           int64     `json:"id" xml:"id"`
	Name         string    `json:"name" xml:"name"`
	Created_date time.Time `json:"created_date" xml:"created_date"`
	Updated_date time.Time `json:"updated_date" xml:"updated_date"`
	Work_id      int64     `json:"work_id" xml:"work_id"`
}
