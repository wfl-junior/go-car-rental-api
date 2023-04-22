package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	BrandRepository "github.com/wfl-junior/go-car-rental-api/repositories/brands"
)

func Index(context *gin.Context) {
	// get the brands
	brands, err := BrandRepository.GetAll()

	// return error response if there is an error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	// return the brands
	context.JSON(http.StatusOK, gin.H{
		"brands": brands,
	})
}
