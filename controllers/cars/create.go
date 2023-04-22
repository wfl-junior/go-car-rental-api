package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wfl-junior/go-car-rental-api/initializers"
	"github.com/wfl-junior/go-car-rental-api/models"
)

func Create(context *gin.Context) {
	// get data from body
	var body CarBody
	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})

		return
	}

	// create and save in database
	car := models.Car{
		BrandId:               body.BrandId,
		Model:                 body.Model,
		PriceInUsd:            body.PriceInUsd,
		HorsePower:            body.HorsePower,
		TorqueInLb:            body.TorqueInLb,
		TopSpeedInKm:          body.TopSpeedInKm,
		AccelerationSpeedInKm: body.AccelerationSpeedInKm,
		WeightInKg:            body.WeightInKg,
	}

	result := initializers.DB.Create(&car)

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

	// return new car
	context.JSON(http.StatusCreated, gin.H{
		"car": car,
	})
}
