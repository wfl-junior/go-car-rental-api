package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wfl-junior/go-car-rental-api/initializers"
	"github.com/wfl-junior/go-car-rental-api/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(context *gin.Context) {
	// get data from body
	var body RegisterBody
	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})

		return
	}

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	// create user and save in database
	user := models.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: string(hash),
	}

	result := initializers.DB.Create(&user)

	// return error response if there is an error
	if result.Error != nil {
		errorMessage := result.Error.Error()
		if strings.Contains(errorMessage, "unique constraint") {
			context.JSON(http.StatusConflict, gin.H{
				"error": "User e-mail already registered",
			})

			return
		}

		context.JSON(http.StatusInternalServerError, gin.H{
			"error": errorMessage,
		})

		return
	}

	// return new user
	context.JSON(http.StatusCreated, gin.H{
		"user": user,
	})
}
