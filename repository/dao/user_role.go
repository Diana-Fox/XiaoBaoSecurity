package dao

import (
	"context"
	"gorm.io/gorm"
)

type UserRole struct {
	Id  int64 `gorm:"primaryKey,autoIncrement"`
	UId int64 `json:"u_id"`
	RId int64 `json:"r_id"`
}

type userRoleDao struct {
	db *gorm.DB
}

func NewUserRoleDao(db *gorm.DB) UserRoleDao {
	return &userRoleDao{
		db: db,
	}
}
func (u *userRoleDao) table() string {
	return "user_roles"
}
func (u *userRoleDao) FindRIdsByUId(ctx context.Context, uId int64) ([]int64, error) {
	var ur []int64
	err := u.db.WithContext(ctx).Debug().Table(u.table()).Select([]string{"r_id"}).Where("u_id=?", uId).Find(&ur).Error
	return ur, err
}
