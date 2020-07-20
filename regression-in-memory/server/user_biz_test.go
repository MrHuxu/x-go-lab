package main

import (
	"context"
	"testing"

	"github.com/MrHuxu/x-go-lab/regression-in-memory/server/models"

	"github.com/stretchr/testify/assert"
)

func TestDBOperations(t *testing.T) {
	assert := assert.New(t)

	biz := newUserBiz()
	ctx := context.Background()

	// case 1: initial count of user should be 2
	list, err := biz.Query(ctx, "")
	assert.Equal(2, len(list))
	assert.Nil(err)

	// case 2: insert a record to table user
	u := &models.User{Name: "baz", Age: 42}
	err = biz.Craete(ctx, u)
	assert.Nil(err)

	// case 3: query table user with age 4 should return 3 records
	list, err = biz.Query(ctx, "4")
	assert.Equal(0, len(list))
	assert.Nil(err)

	// case 4: query table user with age 42 should return the record just created in case 2
	list, err = biz.Query(ctx, "42")
	assert.Equal(1, len(list))
	assert.Equal("baz", list[0].Name)
	assert.Equal(42, list[0].Age)
	assert.Nil(err)
}
