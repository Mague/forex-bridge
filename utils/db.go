package utils

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Query(fn func(*gorm.DB)) {
	var appConfig map[string]string
	appConfig, err := godotenv.Read()

	if err != nil {
		log.Fatal("Error reading .env file")
	}
	dbName := appConfig["DBNAME"]
	if dbName == "" {
		dbName = "db"
	}
	db, err := gorm.Open(sqlite.Open(dbName))
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("Error al obtener la instancia")
	}
	if err != nil {
		panic("failed to connect database")
	}
	fn(db)
	defer sqlDB.Close()
}
