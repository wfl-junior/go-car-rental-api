package main

import (
	"github.com/gin-gonic/gin"
	BrandController "github.com/wfl-junior/go-car-rental-api/controllers/brands"
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

	// brands
	router.GET("/brands", BrandController.Index)
	router.POST("/brands", BrandController.Create)
	router.GET("/brands/:id", BrandController.Show)
	router.PUT("/brands/:id", BrandController.Update)
	router.DELETE("/brands/:id", BrandController.Delete)

	router.Run()
}
