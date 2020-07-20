package main

import (
	"database/sql"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *sql.DB

func initDB() {
	var err error

	db, err = sql.Open("mysql", "root:@/regression")
	if err != nil {
		panic(err)
	}
}
