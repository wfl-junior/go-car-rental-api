package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	RentalRepository "github.com/wfl-junior/go-car-rental-api/repositories/rentals"
)

func Index(context *gin.Context) {
	// get the rentals
	rentals, err := RentalRepository.GetAll()

	// return error response if there is an error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	// return the rentals
	context.JSON(http.StatusOK, gin.H{
		"rentals": rentals,
	})
}
