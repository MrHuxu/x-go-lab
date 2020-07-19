package main

import (
	"io/ioutil"
	"testing"

	"github.com/jinzhu/gorm"
)

func TestMain(m *testing.M) {
	var err error

	// prepare in-memory database and craete table
	db, err = gorm.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	db.CreateTable(&user{})

	// load the content of data.sql to import data
	bytes, err := ioutil.ReadFile("./sqls/data.sql")
	if err != nil {
		panic(err)
	}
	if err := db.Exec(string(bytes)).Error; err != nil {
		panic(err)
	}

	// run unit test cases
	m.Run()
}
