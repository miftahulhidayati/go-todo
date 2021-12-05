package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysqlDB() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/go_gin_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// db.AutoMigrate(user.User{})
	return db
}