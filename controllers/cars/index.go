package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wfl-junior/go-car-rental-api/initializers"
	"github.com/wfl-junior/go-car-rental-api/models"
)



func Index(context *gin.Context) {
	// get the cars
	var cars []models.Car
	result := initializers.DB.Find(&cars)

	// return error response if there is an error
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H {
			"error": result.Error.Error(),
		})

		return
	}

	// return the cars
	context.JSON(http.StatusOK, gin.H {
		"cars": cars,
	})
}