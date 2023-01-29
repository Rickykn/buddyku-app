package main

import (
	"fmt"
	"log"

	"github.com/Rickykn/buddyku-app.git/routers"
	"github.com/joho/godotenv"
)

func main() {

	fmt.Println("Buddyku-app")

	errEnv := godotenv.Load()

	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	// gin.SetMode(gin.DebugMode)

	// server := http.Server{
	// 	Addr:    ":8080",
	// 	Handler: routers.Server(),
	// }

	// go func() {
	// 	if err := server.ListenAndServe(); err != nil {
	// 		log.Printf("Server Listen Error : %s ", err.Error())
	// 	}
	// }()

	// signals := make(chan os.Signal)

	// signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	// <-signals
	// log.Println("Server is Shutdown")

	server := routers.Server()

	err := server.Run()
	if err != nil {
		panic(err)
	}
}
