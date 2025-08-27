package nivasDB

import (
	"database/sql"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error in load the .env File In DB Connection")
	}

}

func initDB(*gorm.DB, *sql.DB) {

}
