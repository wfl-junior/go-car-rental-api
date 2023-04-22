package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wfl-junior/go-car-rental-api/initializers"
	"github.com/wfl-junior/go-car-rental-api/models"
)

type CarBody struct {
	Brand                 string  `binding:"required"`
	Model 								string  `binding:"required"`
	PriceInUsd            float64 `binding:"required"`
	HorsePower            uint16  `binding:"required"`
	TorqueInLb            uint16  `binding:"required"`
	TopSpeedInKm          uint16  `binding:"required"`
	AccelerationSpeedInKm float32 `binding:"required"`
	WeightInKg            uint16  `binding:"required"`
}

func CarsCreate(context *gin.Context) {
	// get data from body
	var body CarBody
	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H {
			"error": err.Error(),
		})

		return
	}

	// create and save in database
	car := models.Car {
		Brand: body.Brand,
		Model: body.Model,
		PriceInUsd: body.PriceInUsd,
		HorsePower: body.HorsePower,
		TorqueInLb: body.TorqueInLb,
		TopSpeedInKm: body.TopSpeedInKm,
		AccelerationSpeedInKm: body.AccelerationSpeedInKm,
		WeightInKg: body.WeightInKg,
	}

	result := initializers.DB.Create(&car)

	// return error response if there is an error
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H {
			"error": result.Error.Error(),
		})

		return
	}

	// return it
	context.JSON(http.StatusCreated, gin.H {
		"car": car,
	})
}