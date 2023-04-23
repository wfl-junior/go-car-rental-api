package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/wfl-junior/go-car-rental-api/models"
)

func GetUserFromContext(context *gin.Context) models.User {
	maybeUser, _ := context.Get("user")
	return maybeUser.(models.User)
}
