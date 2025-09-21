package main

import (
	"food-ordering/config"
	"food-ordering/models"
	"food-ordering/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func InitializeApp() *gin.Engine {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env")
	}

	r := gin.Default()
	db := config.ConnectDatabase()
	
	db.AutoMigrate(&models.User{}, &models.Menu{}, &models.Order{}, &models.Cart{})

	routes.SetupRoutes(r, db)
	return r
}

func main() {
	app := InitializeApp()
	app.Run(":8080")
}