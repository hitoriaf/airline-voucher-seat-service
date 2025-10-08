package main

import (
	"log"
	"os"

	"airline-voucher-seat-service/src/database"
	"airline-voucher-seat-service/src/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading configuration")
	}
	dbPath := os.Getenv("DB_PATH")
	port := os.Getenv("PORT")

	if dbPath == "" || port == "" {
		log.Fatal("PORT or DB_PATH is not set")
	}
	database.InitDB()

	router := gin.Default()
	router.Use(cors.Default())
	// Setup API routes
	routes.SetupRoutes(router)
	err = router.Run()
	if err != nil {
		return
	}
}
