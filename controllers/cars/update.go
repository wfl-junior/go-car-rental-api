package controllers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wfl-junior/go-car-rental-api/initializers"
	"github.com/wfl-junior/go-car-rental-api/models"
	CarRepository "github.com/wfl-junior/go-car-rental-api/repositories/cars"
	"gorm.io/gorm"
)

func Update(context *gin.Context) {
	// get data from body
	var body CarBody
	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})

		return
	}

	// get the id from the path params
	id := context.Param("id")

	// get the car by id
	car, err := CarRepository.GetById(id)

	// return error response if there is an error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusNotFound, gin.H{
				"error": "Car not found",
			})

			return
		}

		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	// update and save in database
	result := initializers.DB.Model(&car).Updates(models.Car{
		BrandId:               body.BrandId,
		Model:                 body.Model,
		RentalPriceDailyInUsd: body.RentalPriceDailyInUsd,
		HorsePower:            body.HorsePower,
		TorqueInLb:            body.TorqueInLb,
		TopSpeedInKm:          body.TopSpeedInKm,
		AccelerationSpeedInKm: body.AccelerationSpeedInKm,
		WeightInKg:            body.WeightInKg,
	})

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

	// return updated car
	context.JSON(http.StatusCreated, gin.H{
		"car": car,
	})
}
