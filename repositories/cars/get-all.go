package repositories

import (
	"github.com/wfl-junior/go-car-rental-api/initializers"
	"github.com/wfl-junior/go-car-rental-api/models"
)

func GetAll() ([]models.Car, error) {
	var cars []models.Car
	err := initializers.DB.Preload("Brand").Find(&cars).Error

	return cars, err
}
