package controllers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wfl-junior/go-car-rental-api/initializers"
	"github.com/wfl-junior/go-car-rental-api/models"
	BrandRepository "github.com/wfl-junior/go-car-rental-api/repositories/brands"
	"gorm.io/gorm"
)

func Update(context *gin.Context) {
	// get data from body
	var body BrandBody
	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})

		return
	}

	// get the id from the path params
	id := context.Param("id")

	// get the brand by id
	brand, err := BrandRepository.GetById(id)

	// return error response if there is an error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusNotFound, gin.H{
				"error": "Brand not found",
			})

			return
		}

		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	// update and save in database
	result := initializers.DB.Model(&brand).Updates(models.Brand{
		Name: body.Name,
	})

	// return error response if there is an error
	if result.Error != nil {
		errorMessage := result.Error.Error()
		if strings.Contains(errorMessage, "unique constraint") {
			context.JSON(http.StatusConflict, gin.H{
				"error": "Brand name already exists",
			})

			return
		}

		context.JSON(http.StatusInternalServerError, gin.H{
			"error": errorMessage,
		})

		return
	}

	// return updated brand
	context.JSON(http.StatusCreated, gin.H{
		"brand": brand,
	})
}
