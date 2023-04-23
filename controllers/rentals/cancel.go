package controllers

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wfl-junior/go-car-rental-api/initializers"
	"github.com/wfl-junior/go-car-rental-api/models"
	RentalRepository "github.com/wfl-junior/go-car-rental-api/repositories/rentals"
	"github.com/wfl-junior/go-car-rental-api/utils"
	"gorm.io/gorm"
)

func Cancel(context *gin.Context) {
	// get id from path params
	id := context.Param("id")

	// get the rental by id
	rental, err := RentalRepository.GetById(id)

	// return error response if there is an error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusNotFound, gin.H{
				"error": "Rental not found",
			})

			return
		}

		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	user := utils.GetUserFromContext(context)

	// return forbidden error if the user trying to cancel the rental is not the owner
	if rental.UserId != user.Id {
		context.JSON(http.StatusForbidden, gin.H{
			"error": "You must be the owner to cancel the rental",
		})

		return
	}

	// return bad request if the rental is already canceled
	if rental.CanceledAt != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "This rental has already been canceled",
		})

		return
	}

	// return bad request if the rental start date has passed
	if rental.StartsAt.Before(time.Now()) {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "You cannot cancel a rental that has already started",
		})

		return
	}

	// update canceled_at and save in database
	now := time.Now()
	result := initializers.DB.Model(&rental).Updates(models.Rental{
		CanceledAt: &now,
	})

	// return error response if there is an error updating
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})

		return
	}

	// return updated rental
	context.JSON(http.StatusOK, gin.H{
		"rental": rental,
	})
}
