package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite3", "./sqlite.db")

	stmt, _ := db.Prepare("CREATE TABLE IF NOT EXISTS users(username, password)")
	stmt.Exec()

	stmt, _ = db.Prepare("INSERT INTO users(username, password) VALUES (?, ?)")
	stmt.Exec("kruztw", nil)

	checkError(err)
	query := "select * from users;"
	rows, err := db.Query(query)
	checkError(err)

	var pw sql.NullString // avoid error "converting NULL to string is unsupported"
	for rows.Next() {
		var username, password string
		err = rows.Scan(&username, &pw)
		checkError(err)

		if pw.Valid {
			password = pw.String
		}

		fmt.Printf("username = %s, password = %s\n", username, password)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}
