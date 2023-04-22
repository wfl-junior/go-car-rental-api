package repositories

import (
	"github.com/wfl-junior/go-car-rental-api/initializers"
	"github.com/wfl-junior/go-car-rental-api/models"
)

func GetById(id string) (models.User, error) {
	var user models.User
	err := initializers.DB.Where("id = ?", id).First(&user).Error

	return user, err
}
