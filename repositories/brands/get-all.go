package repositories

import (
	"github.com/wfl-junior/go-car-rental-api/initializers"
	"github.com/wfl-junior/go-car-rental-api/models"
)

func GetAll() ([]models.Brand, error) {
	var brands []models.Brand
	err := initializers.
		DB.
		Preload("Cars").
		Order("created_at asc").
		Find(&brands).
		Error

	return brands, err
}
