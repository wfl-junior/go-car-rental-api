package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	UserMappers "github.com/wfl-junior/go-car-rental-api/mappers/users"
	"github.com/wfl-junior/go-car-rental-api/models"
)

func Me(context *gin.Context) {
	// get user from request
	maybeUser, _ := context.Get("user")
	user := maybeUser.(models.User)

	// return user and jwt
	context.JSON(http.StatusOK, gin.H{
		"user": UserMappers.ToBaseViewModel(user),
	})
}
