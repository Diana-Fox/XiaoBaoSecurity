package ioc

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13306)/rbac"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
