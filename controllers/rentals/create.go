package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wfl-junior/go-car-rental-api/initializers"
	"github.com/wfl-junior/go-car-rental-api/models"
	RentalRepository "github.com/wfl-junior/go-car-rental-api/repositories/rentals"
	"github.com/wfl-junior/go-car-rental-api/utils"
)

func Create(context *gin.Context) {
	// get data from body
	var body RentalBody
	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})

		return
	}

	// return bad request error if the period is in the past
	if body.StartsAt.Before(time.Now()) || body.EndsAt.Before(time.Now()) {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "The period cannot be in the past",
		})

		return
	}

	// return bad request error if the period end date is before start date
	if body.EndsAt.Before(body.StartsAt) {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "The period end date cannot be before the period start date",
		})

		return
	}

	// return bad request error if the period is less than 1 hour
	if body.EndsAt.Unix()-body.StartsAt.Unix() < int64(time.Hour.Seconds()) {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "The period must be at least 1 hour",
		})

		return
	}

	user := utils.GetUserFromContext(context)

	hasOverlapping, err := RentalRepository.HasOverlappingByCarId(
		body.CarId,
		body.StartsAt,
		body.EndsAt,
	)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	if hasOverlapping {
		context.JSON(http.StatusConflict, gin.H{
			"error": "This car is already rented during the specified period",
		})

		return
	}

	// create and save in database
	rental := models.Rental{
		UserId:   user.Id,
		CarId:    body.CarId,
		StartsAt: body.StartsAt,
		EndsAt:   body.EndsAt,
	}

	result := initializers.DB.Create(&rental)

	// return error response if there is an error
	if result.Error != nil {
		errorMessage := result.Error.Error()
		if strings.Contains(errorMessage, "unique constraint") {
			context.JSON(http.StatusConflict, gin.H{
				"error": "Car model already exists for this brand",
			})

			return
		}

		context.JSON(http.StatusInternalServerError, gin.H{
			"error": errorMessage,
		})

		return
	}

	// return new rental
	context.JSON(http.StatusCreated, gin.H{
		"rental": rental,
	})
}
