package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"

	"github.com/Mague/forex-bridge/models"
	"github.com/Mague/forex-bridge/utils"
)

var router *gin.Engine

func LoadDatabase() {
	utils.Query(func(db *gorm.DB) {
		db.AutoMigrate(
			&models.Operation{},
		)
	})
}
func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

func main() {
	LoadDatabase()
	ConfigRuntime()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	} else {
		fmt.Println("Siempree lee el archivo")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Fatal("$PORT must be set")
	}
	gin.SetMode(gin.DebugMode)

	router := gin.Default()

	router.Use(CORSMiddleware())
	router.Use(gin.Recovery())
	router.LoadHTMLGlob("public/*")
	initalizeRoutes(router)
	router.Run(":" + port)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
