package repositories

import (
	"github.com/wfl-junior/go-car-rental-api/initializers"
	"github.com/wfl-junior/go-car-rental-api/models"
)

func GetById(id string) (models.Brand, error) {
	var brand models.Brand
	err := initializers.DB.Where("id = ?", id).Preload("Cars").Take(&brand).Error

	return brand, err
}
