package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDBOperations(t *testing.T) {
	assert := assert.New(t)

	biz := newUserBiz()
	ctx := context.Background()

	// case 1: initial count of user should be 0
	list, err := biz.Query(ctx, nil)
	assert.Equal(0, len(list))
	assert.Nil(err)

	// case 2: insert a record to table user
	u := user{Name: "foo", Age: 42}
	err = biz.Craete(ctx, u)
	assert.Nil(err)

	// case 3: query table user with age 4 should return empty list
	list, err = biz.Query(ctx, map[string]interface{}{"age": 4})
	assert.Equal(0, len(list))
	assert.Nil(err)

	// case 4: query table user with age 42 should return the record just created in case 2
	list, err = biz.Query(ctx, map[string]interface{}{"age": 42})
	assert.Equal(1, len(list))
	assert.Equal("foo", list[0].Name)
	assert.Equal(42, list[0].Age)
	assert.Nil(err)
}
