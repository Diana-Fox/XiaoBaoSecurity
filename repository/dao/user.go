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
	Id       int64          `json:"id" gorm:"primaryKey,autoIncrement"`
	Email    sql.NullString `json:"email" gorm:"unique"`
	Password string         `json:"password"`
	Ctime    int64          `json:"ctime"`
	Utime    int64          `json:"utime"`
}

func NewUserDao(db *gorm.DB) UserDao {
	return &DefaultUserDao{
		db: db,
	}
}
func (d *DefaultUserDao) table() string {
	return "users"
}
func (d *DefaultUserDao) Insert(ctx context.Context, user User) error {
	now := time.Now().UnixMilli()
	user.Ctime = now
	user.Utime = now
	return d.db.WithContext(ctx).Table(d.table()).Create(&user).Error
}

func (d *DefaultUserDao) FindByEmail(ctx context.Context, email string) (User, error) {
	var u User
	err := d.db.WithContext(ctx).Table(d.table()).Where("email=?", email).Find(&u).Error
	return u, err
}
