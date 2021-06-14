package main

import (
       "fmt"
	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

var db *xorm.Engine


type table1 struct {
	Id   int64
	Col1 string `gorm:"unique;not null";json:"flag"`   // column 開頭要大寫
}

type table2 struct {
	Id   int64
	Col2 string
        Col22 int64
}

func main() {
	db, _ = xorm.NewEngine("sqlite3", "./simple.db")
	defer db.Close()
	db.Sync2(new(table1))
	db.Sync2(new(table2))
	db.Insert(table1{Col1: "I am col1"},)
	db.Insert(table2{Col2: "I am col2", Col22:1337},)

        var t2 table2
        db.Select("Id, Col2").Where("Id = ? AND Col22 = ?", 1, 1337).Get(&t2)

        fmt.Println(t2)
}
