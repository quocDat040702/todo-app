package userstorage

import (
	"context"
	"todo/common"
	"todo/module/user/usermodel"
)

func (s *sqlStore) CreateUser(ctx context.Context, data *usermodel.User) error {
	db := s.db.Begin()

	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
