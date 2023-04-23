package repositories

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/wfl-junior/go-car-rental-api/initializers"
	"github.com/wfl-junior/go-car-rental-api/models"
)

func HasOverlappingByCarId(
	carId uuid.UUID,
	startsAt time.Time,
	endsAt time.Time,
) (bool, error) {
	var exists bool
	err := initializers.
		DB.
		Model(models.Rental{}).
		Select("COUNT(*) > 0").
		Where(
			`car_id = @car_id AND canceled_at IS NULL AND (
				(starts_at < @ends_at) AND (ends_at > @starts_at)
			)`,
			sql.Named("car_id", carId),
			sql.Named("starts_at", startsAt),
			sql.Named("ends_at", endsAt),
		).
		Find(&exists).
		Error

	return exists, err
}
