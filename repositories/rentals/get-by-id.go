package repositories

import (
	"github.com/wfl-junior/go-car-rental-api/initializers"
	"github.com/wfl-junior/go-car-rental-api/models"
)

func GetById(id string) (models.Rental, error) {
	var rental models.Rental
	err := initializers.DB.Where("id = ?", id).Take(&rental).Error

	return rental, err
}
