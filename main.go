package main

import (
	"github.com/gin-gonic/gin"
	CarController "github.com/wfl-junior/go-car-rental-api/controllers/cars"
	"github.com/wfl-junior/go-car-rental-api/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	router := gin.Default()

	// cars
	router.GET("/cars", CarController.Index)
	router.POST("/cars", CarController.Create)
	router.GET("/cars/:id", CarController.Show)
	router.PUT("/cars/:id", CarController.Update)
	router.DELETE("/cars/:id", CarController.Delete)

	router.Run()
}