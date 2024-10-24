package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitMySQL() (err error) {
	dsn := "root:yuchao123@tcp(127.0.0.1:3306)/todolist?charset=utf8mb4"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}
