package main

import (
	"context"

	"github.com/MrHuxu/x-go-lab/regression-in-memory/server/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type userBiz interface {
	Query(ctx context.Context, ageCond string) ([]*models.User, error)
	Craete(ctx context.Context, data *models.User) error
}

func newUserBiz() userBiz {
	return &userBizImpl{}
}

type userBizImpl struct {
}

func (i *userBizImpl) Query(ctx context.Context, ageCond string) ([]*models.User, error) {
	var userSlice models.UserSlice
	var err error

	if ageCond == "" {
		userSlice, err = models.Users(
			Select("id", "name", "age"),
		).All(ctx, db)
	} else {
		userSlice, err = models.Users(
			Select("id", "name", "age"),
			Where("age = ?", ageCond),
		).All(ctx, db)
	}

	return ([]*models.User)(userSlice), err
}

func (i *userBizImpl) Craete(ctx context.Context, data *models.User) (err error) {
	return data.Insert(ctx, db, boil.Columns{
		Cols: []string{"name", "age", "created_at", "updated_at"},
	})
}
