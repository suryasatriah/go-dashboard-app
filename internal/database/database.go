package database

import (
	"fmt"
	"log"

	"github.com/surysatriah/go-dashboard-app/internal/model"
	"github.com/surysatriah/go-dashboard-app/pkg"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbHost     = pkg.GetDotEnvVariable("DB_HOST")
	dbUser     = pkg.GetDotEnvVariable("DB_USER")
	dbPassword = pkg.GetDotEnvVariable("DB_PASSWORD")
	dbName     = pkg.GetDotEnvVariable("DB_NAME")
	dbPort     = pkg.GetDotEnvVariable("DB_PORT")
	db         *gorm.DB
	err        error
)

func ConnectDatabase() {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database, error: %v", err)
	}
	log.Println("Successfuly connected to database.")

	migrateModel()

}

func GetDatabase() *gorm.DB {
	return db
}

func migrateModel() {
	db.Debug().AutoMigrate(model.User{}, model.Payload{})
	if err != nil {
		log.Fatalf("Failed to migrate models, error: %v", err)
	}
}
