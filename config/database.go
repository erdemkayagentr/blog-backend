package config

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() *gorm.DB {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_HOST")), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
