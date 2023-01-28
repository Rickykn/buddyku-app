package main

import (
	"fmt"
	"log"

	"github.com/Rickykn/buddyku-app.git/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	fmt.Println("Buddyku-app")

	errEnv := godotenv.Load()

	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	errConnect := database.Connect()

	if errConnect != nil {
		panic(errConnect)
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	err := r.Run()

	if err != nil {
		panic(err)
	}
}
