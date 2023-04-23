package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	RentalRepository "github.com/wfl-junior/go-car-rental-api/repositories/rentals"
)

func Index(context *gin.Context) {
	// get query params
	carId := context.Query("CarId")
	startsAt := context.Query("StartsAt")
	endsAt := context.Query("EndsAt")

	// get the rentals
	rentals, err := RentalRepository.GetAll(RentalRepository.GetAllParams{
		CarId:    carId,
		StartsAt: startsAt,
		EndsAt:   endsAt,
	})

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
