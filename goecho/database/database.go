package database

import (
	"Pea-backend/functions"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

func Call_DB() *sql.DB {
	DB_URL := functions.GoDotEnv("DB_URL")
	db, err := sql.Open("postgres", DB_URL)
	if err != nil {
		// fmt.Fprint(os.Stderr, "Unable to connect database : %v\n", err)
		fmt.Printf("Unable to connect database : %v\n", err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("database connection error : %v\n", err)
	}
	// defer conn.Close(context.Background())
	return db
}

func Init_table() {
	db := Call_DB()
	// ! role_info
	_, err := db.Exec(
		`
		CREATE TABLE IF NOT EXISTS role_info(
			id SERIAL NOT NULL,
			name varchar(30) NOT NULL,
			PRIMARY KEY (id)
		)
	`)
	if err != nil {
		fmt.Printf("Error from create role table :  %v\n", err)
	}

	// ! work_info
	_, err = db.Exec(
		`
		CREATE TABLE IF NOT EXISTS work_info(
			id SERIAL NOT NULL,
			name varchar(100) NOT NULL,
			created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        	updated_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (id)
		)
	`)
	if err != nil {
		fmt.Printf("Error from create work_info table :  %v\n", err)
	}

	// ! group_info
	_, err = db.Exec(
		`
		CREATE TABLE IF NOT EXISTS group_info(
			id SERIAL NOT NULL,
			name varchar(100) NOT NULL,
			created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        	updated_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			work_id int,
			PRIMARY KEY (id),
			FOREIGN KEY (work_id) REFERENCES work_info(id)
		)
	`)
	if err != nil {
		fmt.Printf("Error from create group_info table :  %v\n", err)
	}

	// ! user_info
	_, err = db.Exec(
		`
		CREATE TABLE IF NOT EXISTS user_info(
			id SERIAL NOT NULL,
			username varchar(255) NOT NULL,
			password varchar(30) NOT NULL,
			group_id int, 
			role_id int, 
			PRIMARY KEY (username),
			FOREIGN KEY (role_id) REFERENCES role_info(id),
			FOREIGN KEY (group_id) REFERENCES group_info(id)
		)
	`)
	if err != nil {
		fmt.Printf("Error from create user table :  %v\n", err)
	}

	defer db.Close()
}

func Init_data() {
	db := Call_DB()
	role := strings.Split(functions.GoDotEnv("ROLE"), ", ")
	sql_stm := `INSERT INTO role_info(id, name) VALUES %s ON CONFLICT DO NOTHING`
	sql_values := []string{}
	for index, ele := range role {
		sql_values = append(sql_values, "("+strconv.Itoa(index)+", '"+ele+"')")
	}
	sql_stm = fmt.Sprintf(sql_stm, strings.Join(sql_values, ","))
	sql_prepare, err := db.Prepare(sql_stm)
	if err != nil {
		fmt.Printf("Error from init PREPARE insert role_values : %v\n", err)
	}
	_, err = sql_prepare.Exec()
	if err != nil {
		fmt.Printf("Error from init EXEC insert role_values : %v\n", err)
	}

	work := strings.Split(functions.GoDotEnv("WORK"), ", ")
	sql_stm = `INSERT INTO work_info(id, name) VALUES %s ON CONFLICT DO NOTHING`
	sql_values = []string{}
	for index, ele := range work {
		sql_values = append(sql_values, "("+strconv.Itoa(index)+", '"+ele+"')")
	}
	sql_stm = fmt.Sprintf(sql_stm, strings.Join(sql_values, ", "))
	sql_prepare, err = db.Prepare(sql_stm)
	if err != nil {
		fmt.Printf("Error from init PREPARE insert work_values : %v\n", err)
	}
	_, err = sql_prepare.Exec()
	if err != nil {
		fmt.Printf("Error from init EXEC insert work_values : %v\n", err)
	}

	admin := strings.Split(functions.GoDotEnv("ADMIN"), ", ")
	sql_stm = `INSERT INTO user_info(username, password, role_id) VALUES ('` + admin[0] + "', '" + admin[1] + "', 2" + `) ON CONFLICT DO NOTHING`
	sql_prepare, err = db.Prepare(sql_stm)
	if err != nil {
		fmt.Printf("Error from init PREPARE insert ADMIN: %v\n", err)
	}
	_, err = sql_prepare.Exec()
	if err != nil {
		fmt.Printf("Error from init EXEC insert ADMIN: %v\n", err)
	}
	user := strings.Split(functions.GoDotEnv("USER"), ", ")
	sql_stm = `INSERT INTO user_info(username, password, role_id) VALUES ('` + user[0] + "', '" + user[1] + "', 1" + `) ON CONFLICT DO NOTHING`
	sql_prepare, err = db.Prepare(sql_stm)
	if err != nil {
		fmt.Printf("Error from init PREPARE insert USER : %v\n", err)
	}
	_, err = sql_prepare.Exec()
	if err != nil {
		fmt.Printf("Error from init EXEC insert USER : %v\n", err)
	}

}
