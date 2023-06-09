package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	CarRepository "github.com/wfl-junior/go-car-rental-api/repositories/cars"
)

func Index(context *gin.Context) {
	// get the cars
	cars, err := CarRepository.GetAll()

	// return error response if there is an error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	// return the cars
	context.JSON(http.StatusOK, gin.H{
		"cars": cars,
	})
}
