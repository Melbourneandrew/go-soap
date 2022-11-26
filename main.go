package main

import (
	"fmt"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/melbourneandrew/go-soap/m/v2/controllers"
	"github.com/melbourneandrew/go-soap/m/v2/models"
)

func main() {
	// Load dotenv
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}
	models.ConnectDatabase()
	// initalize gin router
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"m": "Hello, Soap!",
		})
	})

	auth := r.Group("/auth")
	auth.POST("/login", controllers.Login)
	auth.POST("/signup", controllers.Signup)

	err := r.Run()
	if err != nil {
		fmt.Println("Server run failed")
		return
	}
}
