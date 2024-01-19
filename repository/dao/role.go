package dao

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type Role struct {
	ParentId int64  //父角色，一次性选中底下所有
	Id       int64  `gorm:"primaryKey,autoIncrement"`
	Name     string //角色名称
	Ctime    int64
	Utime    int64
}
type roleDao struct {
	db *gorm.DB
}

func (r *roleDao) FindByIds(ctx context.Context, ids []int64) ([]Role, error) {
	var rs []Role
	err := r.db.WithContext(ctx).Find(&rs).Error
	return rs, err
}

func NewRoleDao(db *gorm.DB) RoleDao {
	return &roleDao{
		db: db,
	}
}
func (r *roleDao) Insert(ctx context.Context, role Role) error {
	now := time.Now().UnixMilli()
	role.Ctime = now
	role.Utime = now
	err := r.db.WithContext(ctx).Error
	return err
}

func (r *roleDao) FindAll(ctx context.Context) (Role, error) {
	//TODO implement me
	//查出全部然后Repository层去组装数据，一般是后台配置用，暂时不管
	panic("implement me")
}
