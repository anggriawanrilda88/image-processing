package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/image-processing/src/interface/rest"
	"github.com/joho/godotenv"
)

func init() {
	//To load our environmental variables.
	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}
}

func main() {
	gin.SetMode(os.Getenv("GIN_MODE"))

	// set routes
	r := gin.Default()

	// set routes all interface
	rest.NewRoutes(r)

	app_port := os.Getenv("PORT")
	if app_port == "" {
		app_port = os.Getenv("API_PORT") //localhost
	}
	log.Fatal(r.Run(":" + app_port))
}
