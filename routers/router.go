package routers

import (
	"github.com/Rickykn/buddyku-app.git/database"
	"github.com/Rickykn/buddyku-app.git/handlers"
	"github.com/Rickykn/buddyku-app.git/repositories"
	"github.com/Rickykn/buddyku-app.git/services"
	"github.com/gin-gonic/gin"
)

func Server() *gin.Engine {
	engine := gin.New()
	errConnect := database.Connect()

	ur := repositories.NewUserRepository(&repositories.URConfig{
		DB: database.Get(),
	})

	us := services.NewUserService(&services.USConfig{
		UserRepository: ur,
	})

	h := handlers.New(&handlers.HandlerConfig{

		UserService: us,
	})

	users := engine.Group("/users")
	{

		users.POST("/register", h.RegisterUser)
	}

	if errConnect != nil {
		panic(errConnect)
	}
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return engine
}
