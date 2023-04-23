package models

import (
	"time"

	"github.com/google/uuid"
)

type Car struct {
	Id                    uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	BrandId               uuid.UUID `gorm:"type:uuid;index:unique_brand_model,unique;notNull"`
	CreatedAt             time.Time `gorm:"notNull"`
	UpdatedAt             time.Time `gorm:"notNull"`
	Model                 string    `gorm:"type:varchar(255);index:unique_brand_model,unique;notNull"`
	RentalPriceDailyInUsd float64   `gorm:"notNull"`
	HorsePower            uint16    `gorm:"notNull"`
	TorqueInLb            float32   `gorm:"notNull"`
	TopSpeedInKm          uint16    `gorm:"notNull"`
	AccelerationSpeedInKm float32   `gorm:"notNull"`
	WeightInKg            uint16    `gorm:"notNull"`

	Brand   *Brand
	Rentals []Rental
}
