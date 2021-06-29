package database

import (
	"backend/functions"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Call_DB() *sql.DB {
	DB_URL := `postgres://` + functions.GoDotEnv("POSTGRES_USER") +
		`:` + functions.GoDotEnv("POSTGRES_PASSWORD") + `@localhost:5432/` +
		functions.GoDotEnv("POSTGRES_DB") + `sslmode=disable`
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
		CREATE TABLE IF NOT EXISTS users(
			id SERIAL NOT NULL,
			username varchar(30) NOT NULL,
			password varchar(30) NOT NULL,
			PRIMARY KEY (name)
		)
	`)
	if err != nil {
		fmt.Printf("Error from create users table :  %v\n", err)
	}

	defer db.Close()
}

func Init_data() {
	// db := Call_DB()
	// role := strings.Split(functions.GoDotEnv("ROLE"), ", ")
	// sql_stm := `INSERT INTO role_info(id, name) VALUES %s ON CONFLICT DO NOTHING`
	// sql_values := []string{}
	// for index, ele := range role {
	// 	sql_values = append(sql_values, "("+strconv.Itoa(index)+", '"+ele+"')")
	// }
	// sql_stm = fmt.Sprintf(sql_stm, strings.Join(sql_values, ","))
	// sql_prepare, err := db.Prepare(sql_stm)
	// if err != nil {
	// 	fmt.Printf("Error from init PREPARE insert role_values : %v\n", err)
	// }
	// _, err = sql_prepare.Exec()
	// if err != nil {
	// 	fmt.Printf("Error from init EXEC insert role_values : %v\n", err)
	// }

	// work := strings.Split(functions.GoDotEnv("WORK"), ", ")
	// sql_stm = `INSERT INTO work_info(id, name) VALUES %s ON CONFLICT DO NOTHING`
	// sql_values = []string{}
	// for index, ele := range work {
	// 	sql_values = append(sql_values, "("+strconv.Itoa(index)+", '"+ele+"')")
	// }
	// sql_stm = fmt.Sprintf(sql_stm, strings.Join(sql_values, ", "))
	// sql_prepare, err = db.Prepare(sql_stm)
	// if err != nil {
	// 	fmt.Printf("Error from init PREPARE insert work_values : %v\n", err)
	// }
	// _, err = sql_prepare.Exec()
	// if err != nil {
	// 	fmt.Printf("Error from init EXEC insert work_values : %v\n", err)
	// }

	// admin := strings.Split(functions.GoDotEnv("ADMIN"), ", ")
	// sql_stm = `INSERT INTO user_info(username, password, role_id) VALUES ('` + admin[0] + "', '" + admin[1] + "', 2" + `) ON CONFLICT DO NOTHING`
	// sql_prepare, err = db.Prepare(sql_stm)
	// if err != nil {
	// 	fmt.Printf("Error from init PREPARE insert ADMIN: %v\n", err)
	// }
	// _, err = sql_prepare.Exec()
	// if err != nil {
	// 	fmt.Printf("Error from init EXEC insert ADMIN: %v\n", err)
	// }
	// user := strings.Split(functions.GoDotEnv("USER"), ", ")
	// sql_stm = `INSERT INTO user_info(username, password, role_id) VALUES ('` + user[0] + "', '" + user[1] + "', 1" + `) ON CONFLICT DO NOTHING`
	// sql_prepare, err = db.Prepare(sql_stm)
	// if err != nil {
	// 	fmt.Printf("Error from init PREPARE insert USER : %v\n", err)
	// }
	// _, err = sql_prepare.Exec()
	// if err != nil {
	// 	fmt.Printf("Error from init EXEC insert USER : %v\n", err)
	// }

}
