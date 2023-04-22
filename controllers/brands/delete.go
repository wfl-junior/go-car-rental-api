package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wfl-junior/go-car-rental-api/initializers"
	BrandRepository "github.com/wfl-junior/go-car-rental-api/repositories/brands"
	"gorm.io/gorm"
)

func Delete(context *gin.Context) {
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

	// delete the brand
	result := initializers.DB.Delete(&brand)

	// return error response if there is an error
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})

		return
	}

	// return no content
	context.Status(http.StatusNoContent)
}
