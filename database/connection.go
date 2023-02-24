package database

import (
	"github.com/pavanish/go-react-jwt/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=psql dbname=epic_database port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("could not connect to db")
	}
	DB = db
	db.AutoMigrate(&models.User{})
}
