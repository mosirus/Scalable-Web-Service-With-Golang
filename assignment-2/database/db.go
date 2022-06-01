package database

import (
	"assignment-2/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DB_HOST  = "localhost"
	DB_PORT  = "5432"
	DB_USER  = "monang"
	DB_PASS  = "monang"
	DB_NAME  = "assignment2"
	APP_PORT = ":8888"
)

var (
	db  *gorm.DB
	err error
)

func ConnectDB() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Connecting to database is Error: ", err)
	}
	db.AutoMigrate(models.Order{}, models.Item{})
}

func GetDB() *gorm.DB {
	fmt.Println("Connecting to database is succes")
	return db
}
