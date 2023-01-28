package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitmySQL() (err error) {
	dsn := "root:111111@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return
}
