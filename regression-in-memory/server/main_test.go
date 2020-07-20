package main

import (
	"context"
	"database/sql"
	"io/ioutil"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/volatiletech/sqlboiler/v4/queries"
)

func TestMain(m *testing.M) {
	var err error

	// prepare in-memory database and craete table
	db, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	schema, err := ioutil.ReadFile("./sqls/sqlite3-schema.sql")
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadFile("./sqls/data.sql")
	if err != nil {
		panic(err)
	}
	queries.Raw(string(schema)).Bind(ctx, db, &struct{}{})
	queries.Raw(string(data)).Bind(ctx, db, &struct{}{})

	// run unit test cases
	m.Run()
}
