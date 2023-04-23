package repositories

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/wfl-junior/go-car-rental-api/initializers"
	"github.com/wfl-junior/go-car-rental-api/models"
)

func GetOverlappingByCarId(
	carId uuid.UUID,
	startsAt time.Time,
	endsAt time.Time,
) (models.Rental, error) {
	var rental models.Rental
	err := initializers.DB.Where(
		`car_id = @car_id AND (
			(starts_at < @ends_at) AND (ends_at > @starts_at)
		)`,
		sql.Named("car_id", carId),
		sql.Named("starts_at", startsAt),
		sql.Named("ends_at", endsAt),
	).Take(&rental).Error

	return rental, err
}
