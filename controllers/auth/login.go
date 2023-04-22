package controllers

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	UserMappers "github.com/wfl-junior/go-car-rental-api/mappers/users"
	UserRepository "github.com/wfl-junior/go-car-rental-api/repositories/users"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(context *gin.Context) {
	// get data from body
	var body LoginBody
	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})

		return
	}

	// find user with email from body
	user, err := UserRepository.GetByEmail(body.Email)

	// return invalid credentials error if user is not found
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid credentials",
			})

			return
		}

		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	// compare passwords and return invalid credentials if they are not the same
	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(body.Password),
	); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid credentials",
		})

		return
	}

	// generate jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	secret := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	// return user and jwt
	context.JSON(http.StatusOK, gin.H{
		"accessToken": tokenString,
		"user":        UserMappers.ToBaseViewModel(user),
	})
}
