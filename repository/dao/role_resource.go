package dao

import (
	"context"
	"gorm.io/gorm"
)

type RoleResource struct {
	Id   int64 `json:"id" gorm:"primaryKey,autoIncrement"`
	RId  int64 `json:"r_id"`
	RSId int64 `json:"rs_id"`
}
type roleResourceDao struct {
	db *gorm.DB
}

func NewRoleResourceDao(db *gorm.DB) RoleResourceDao {
	return &roleResourceDao{
		db: db,
	}
}
func (r *roleResourceDao) table() string {
	return "role_resources"
}

// 通过角色查所有权限
func (r roleResourceDao) FindResourceByRoles(ctx context.Context, rIds []int64) ([]int64, error) {
	var rr []int64
	err := r.db.WithContext(ctx).Table("role_resources").Select([]string{"rs_id"}).Where("r_id in (?)", rIds).Find(&rr).Error
	return rr, err
}
