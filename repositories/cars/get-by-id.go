package repositories

import (
	"github.com/wfl-junior/go-car-rental-api/initializers"
	"github.com/wfl-junior/go-car-rental-api/models"
)

func GetById(id string) (models.Car, error) {
	var car models.Car
	err := initializers.DB.Where("id = ?", id).Preload("Brand").First(&car).Error

	return car, err
}
