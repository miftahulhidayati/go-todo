package database

import (
	"fmt"
	"log"
	"os"

	"github.com/miftahulhidayati/go-todo/todo"
	"github.com/miftahulhidayati/go-todo/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitMysqlDB() *gorm.DB {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbName, dbPort)
	fmt.Println(dbUser, dbPass, dbHost, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})

	if err != nil {
		log.Panic("Database error", err.Error())
	}
	db.AutoMigrate(user.User{})
	db.AutoMigrate(todo.Todo{})
	return db
}
