package database

import (
	"fmt"
	"log"

	"github.com/go-pg/pg/v10"
	_ "github.com/lib/pq"
	// "github.com/go-pg/pg/v10/orm"
)

var DB *pg.DB

func ConnectingDatabase(user string, password string, port string, db_name string) {
	DB_URL := fmt.Sprintf("%s:%s", "localhost", port)
	options := &pg.Options{
		User:     user,
		Password: password,
		Addr:     DB_URL,
		Database: db_name,
		PoolSize: 50,
	}
	conn := pg.Connect(options)
	if conn == nil {
		log.Println("Unable to connect database")
	}
	DB = conn
}
