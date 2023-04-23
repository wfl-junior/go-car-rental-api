package repositories

import (
	"github.com/google/uuid"
	"github.com/wfl-junior/go-car-rental-api/initializers"
	"github.com/wfl-junior/go-car-rental-api/models"
)

func GetAllByUserId(userId uuid.UUID) ([]models.Rental, error) {
	var rentals []models.Rental
	err := initializers.
		DB.
		Preload("Car").
		Preload("Car.Brand").
		Order("created_at asc").
		Where("user_id = ?", userId).
		Find(&rentals).
		Error

	return rentals, err
}
