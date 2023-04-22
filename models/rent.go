package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Rental struct {
	Id         uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserId     uuid.UUID `gorm:"type:uuid;notNull"`
	CarId      uuid.UUID `gorm:"type:uuid;notNull"`
	CreatedAt  time.Time `gorm:"notNull"`
	UpdatedAt  time.Time `gorm:"notNull"`
	StartsAt   time.Time `gorm:"notNull"`
	EndsAt     time.Time `gorm:"notNull"`
	CanceledAt sql.NullTime

	User User
	Car  Car
}
