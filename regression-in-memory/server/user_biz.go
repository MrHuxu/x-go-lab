package main

import (
	"context"
)

type user struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func (user) TableName() string {
	return "user"
}

type userBiz interface {
	Query(ctx context.Context, cond map[string]interface{}) ([]*user, error)
	Craete(ctx context.Context, data user) error
}

func newUserBiz() userBiz {
	return &userBizImpl{}
}

type userBizImpl struct {
}

func (i *userBizImpl) Query(ctx context.Context, cond map[string]interface{}) (result []*user, err error) {
	err = db.Where(cond).Find(&result).Error
	return
}

func (i *userBizImpl) Craete(ctx context.Context, data user) (err error) {
	err = db.Create(&data).Error
	return
}
