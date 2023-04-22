package main

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/wfl-junior/go-car-rental-api/controllers/cars"
	"github.com/wfl-junior/go-car-rental-api/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	router := gin.Default()

	// cars
	router.POST("/cars", controllers.CarsCreate)

	router.Run()
}