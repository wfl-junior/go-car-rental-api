package repositories

import (
	"github.com/wfl-junior/go-car-rental-api/initializers"
	"github.com/wfl-junior/go-car-rental-api/models"
)

func GetAll() ([]models.Rental, error) {
	var rentals []models.Rental
	err := initializers.DB.Preload("Car").Preload("Car.Brand").Order("created_at asc").Find(&rentals).Error

	return rentals, err
}
