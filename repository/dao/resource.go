package dao

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type Resource struct {
	ParentId         int64
	Id               int64  `gorm:"primaryKey,autoIncrement"`
	name             string //权限名称
	Url              string //路径
	IsDynamicRouting byte   //是否是动态路由
	DynamicUrl       string //动态路由的表达式，存储的时候对这个做处理
	Level            int    //路由的层级，主要为了匹配动态路由使用
	Ctime            int64
	Utime            int64
}

type resourceDao struct {
	db *gorm.DB
}

func NewResourceDao(db *gorm.DB) ResourceDao {
	return &resourceDao{
		db: db,
	}
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
	err := r.db.WithContext(ctx).Select([]string{"Url"}).Where("id in (?)", ids).Find(&rs).Error
	return err, rs
}
