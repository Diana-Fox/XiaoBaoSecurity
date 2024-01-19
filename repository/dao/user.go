package dao

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type DefaultUserDao struct {
	db *gorm.DB
}

type User struct {
	Id       int64          `gorm:"primaryKey,autoIncrement"`
	Email    sql.NullString `gorm:"unique"`
	Password string
	Ctime    int64
	Utime    int64
}

func NewUserDao(db *gorm.DB) UserDao {
	return &DefaultUserDao{
		db: db,
	}
}

func (d *DefaultUserDao) Insert(ctx context.Context, user User) error {
	now := time.Now().UnixMilli()
	user.Ctime = now
	user.Utime = now
	return d.db.WithContext(ctx).Create(&user).Error
}

func (d *DefaultUserDao) FindByEmail(ctx context.Context, email string) (User, error) {
	var u User
	err := d.db.WithContext(ctx).Where("email=?", email).Find(&u).Error
	return u, err
}
