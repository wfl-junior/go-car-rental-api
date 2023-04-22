package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wfl-junior/go-car-rental-api/initializers"
	CarRepository "github.com/wfl-junior/go-car-rental-api/repositories/cars"
	"gorm.io/gorm"
)

func Delete(context *gin.Context) {
	// get the id from the path params
	id := context.Param("id")

	// get the car by id
	car, err := CarRepository.GetById(id)

	// return error response if there is an error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusNotFound, gin.H {
				"error": "Car not found",
			})
	
			return
		}

		context.JSON(http.StatusInternalServerError, gin.H {
			"error": err.Error(),
		})

		return
	}

	// delete the car
	result := initializers.DB.Delete(&car)

	// return error response if there is an error
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H {
			"error": result.Error.Error(),
		})

		return
	}

	// return no content
	context.Status(http.StatusNoContent)
}