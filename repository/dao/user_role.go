package dao

import (
	"context"
	"gorm.io/gorm"
)

type UserRole struct {
	Id  int64 `gorm:"primaryKey,autoIncrement"`
	UId int64
	RId int64
}

type userRoleDao struct {
	db *gorm.DB
}

func NewUserRoleDao(db *gorm.DB) UserRoleDao {
	return &userRoleDao{
		db: db,
	}
}

func (u *userRoleDao) FindRIdsByUId(ctx context.Context, uId int64) ([]int64, error) {
	var ur []int64
	err := u.db.WithContext(ctx).Table("user_roles").Select([]string{"RId"}).Where("UId=?", uId).Find(&ur).Error
	return ur, err
}
