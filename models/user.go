package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time `gorm:"notNull"`
	UpdatedAt time.Time `gorm:"notNull"`
	Name      string    `gorm:"type:varchar(255);notNull"`
	Email     string    `gorm:"type:varchar(255);uniqueIndex;notNull"`
	Password  string    `gorm:"type:varchar(255);notNull"`

	Rentals []Rental
}
