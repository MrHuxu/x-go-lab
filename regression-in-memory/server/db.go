package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func initDB() {
	var err error
	db, err = gorm.Open("mysql", "root:@/regression?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
}
