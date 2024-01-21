package dao

import (
	"context"
	"gorm.io/gorm"
)

type RoleResource struct {
	Id   int64 `gorm:"primaryKey,autoIncrement"`
	RId  int64
	RSId int64
}
type roleResourceDao struct {
	db *gorm.DB
}

func NewRoleResourceDao(db *gorm.DB) RoleResourceDao {
	return &roleResourceDao{
		db: db,
	}
}

// 通过角色查所有权限
func (r roleResourceDao) FindResourceByRoles(ctx context.Context, rIds []int64) ([]int64, error) {
	var rr []int64
	err := r.db.WithContext(ctx).Table("role_resources").Select([]string{"RSId"}).Where("RId in (?)", rIds).Find(&rr).Error
	return rr, err
}
