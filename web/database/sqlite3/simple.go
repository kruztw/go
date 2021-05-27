package main

import (
    "fmt"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {
    var err error
    db, err = sql.Open("sqlite3", "./sqlite.db")

    checkError(err)
    query := "select * from users;"
    rows, err := db.Query(query)
    checkError(err)

    for rows.Next() {
        var username, password string
        err = rows.Scan(&username, &password)
        checkError(err)
        fmt.Printf("username = %s, password = %s\n", username, password)
    }
}

func checkError(err error) {
    if err != nil {
        fmt.Println(err.Error())
        panic(err)
    }
}
