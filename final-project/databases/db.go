package databases

import (
	"final-project/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbHost     = "localhost"
	dbUser     = "monang"
	dbPassword = "monang"
	dbPort     = "5432"
	dbName     = "finalprojectdb"
	db         *gorm.DB
	err        error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPassword, dbPort, dbName)
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database: ", err)
		return
	}

	fmt.Println("Success for connecting to databases")
	db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.SocialMedia{}, models.Comment{})

}

func GetDB() *gorm.DB {
	return db
}
