package main

import (
	"github.com/gin-gonic/gin"
	AuthController "github.com/wfl-junior/go-car-rental-api/controllers/auth"
	BrandController "github.com/wfl-junior/go-car-rental-api/controllers/brands"
	CarController "github.com/wfl-junior/go-car-rental-api/controllers/cars"
	RentalController "github.com/wfl-junior/go-car-rental-api/controllers/rentals"
	"github.com/wfl-junior/go-car-rental-api/initializers"
	"github.com/wfl-junior/go-car-rental-api/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	router := gin.Default()

	// brands
	router.GET("/brands", BrandController.Index)
	router.POST("/brands", BrandController.Create)
	router.GET("/brands/:id", BrandController.Show)
	router.PUT("/brands/:id", BrandController.Update)
	router.DELETE("/brands/:id", BrandController.Delete)

	// cars
	router.GET("/cars", CarController.Index)
	router.POST("/cars", CarController.Create)
	router.GET("/cars/:id", CarController.Show)
	router.PUT("/cars/:id", CarController.Update)
	router.DELETE("/cars/:id", CarController.Delete)

	// auth
	router.POST("/auth/register", AuthController.Register)
	router.POST("/auth/login", AuthController.Login)
	router.GET("/auth/me", middleware.RequireAuth, AuthController.Me)
	router.GET(
		"/auth/me/rentals",
		middleware.RequireAuth,
		AuthController.MyRentals,
	)

	// rentals
	router.GET("/rentals", RentalController.Index)
	router.POST("/rentals", middleware.RequireAuth, RentalController.Create)
	router.PATCH(
		"/rentals/:id/cancel",
		middleware.RequireAuth,
		RentalController.Cancel,
	)

	router.Run()
}
