package dao

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type Resource struct {
	ParentId int64  `json:"parentId"`
	Id       int64  `json:"id" gorm:"primaryKey,autoIncrement"`
	name     string `json:"name"` //权限名称
	Url      string `json:"url"`  //路径
	Ctime    int64  `json:"ctime"`
	Utime    int64  `json:"utime"`
}

type resourceDao struct {
	db *gorm.DB
}

func NewResourceDao(db *gorm.DB) ResourceDao {
	return &resourceDao{
		db: db,
	}
}
func (r *resourceDao) table() string {
	return "resources"
}
func (r *resourceDao) Insert(ctx context.Context, resource Resource) error {
	now := time.Now().UnixMilli()
	resource.Ctime = now
	resource.Utime = now
	err := r.db.WithContext(ctx).Create(resource).Error
	return err
}

func (r *resourceDao) FindAll(ctx context.Context) (error, Resource) {
	//TODO implement me
	panic("implement me")
}

func (r *resourceDao) FindByIds(ctx context.Context, ids []int64) (error, []Resource) {
	var rs []Resource
	err := r.db.WithContext(ctx).Debug().Table("resources").Where("id in (?)", ids).Find(&rs).Error
	return err, rs
}
