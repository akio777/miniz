package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectingDatabase(user string, password string, port string, db_name string) {
	DB_URL := `postgres://` + user +
		`:` + password + `@` + `database:5434` + `/` +
		db_name + `?sslmode=disable`
	db, err := sql.Open("postgres", DB_URL)
	if err != nil {
		log.Printf("Unable to connect database : %v\n", err)
	}
	err = db.Ping()
	if err != nil {
		log.Printf("database connection error : %v\n", err)
		// 	panic(err)
	}
	DB = db
}
