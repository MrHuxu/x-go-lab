package main

import (
	"testing"

	"github.com/jinzhu/gorm"
)

func TestMain(m *testing.M) {
	// prepare in-memory database and craete table
	db, _ = gorm.Open("sqlite3", ":memory:")
	db.CreateTable(&user{})

	// run unit test cases
	m.Run()
}
