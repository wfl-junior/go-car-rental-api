package repositories

import (
	"github.com/wfl-junior/go-car-rental-api/initializers"
	"github.com/wfl-junior/go-car-rental-api/models"
)

func GetByEmail(email string) (models.User, error) {
	var user models.User
	err := initializers.DB.Where("email = ?", email).Take(&user).Error

	return user, err
}
