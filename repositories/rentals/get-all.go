package repositories

import (
	"github.com/wfl-junior/go-car-rental-api/initializers"
	"github.com/wfl-junior/go-car-rental-api/models"
)

type GetAllParams struct {
	UserId   string
	CarId    string
	StartsAt string
	EndsAt   string
}

func GetAll(params GetAllParams) ([]models.Rental, error) {
	var rentals []models.Rental
	query := initializers.
		DB.
		Preload("Car").
		Preload("Car.Brand").
		Order("starts_at asc")

	if params.UserId != "" {
		query = query.Where("user_id = ?", params.UserId)
	}

	if params.CarId != "" {
		query = query.Where("car_id = ?", params.CarId)
	}

	if params.StartsAt != "" {
		query = query.Where("starts_at >= ?", params.StartsAt)
	}

	if params.EndsAt != "" {
		query = query.Where("ends_at <= ?", params.EndsAt)
	}

	err := query.
		Find(&rentals).
		Error

	return rentals, err
}
