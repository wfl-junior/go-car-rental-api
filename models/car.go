package models

import (
	"time"

	"github.com/google/uuid"
)

type Car struct {
	Id                    uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	BrandId               uuid.UUID `gorm:"type:uuid;index:unique_brand_model,unique"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
	Model                 string `gorm:"type:varchar(255);index:unique_brand_model,unique"`
	PriceInUsd            float64
	HorsePower            uint16
	TorqueInLb            float32
	TopSpeedInKm          uint16
	AccelerationSpeedInKm float32
	WeightInKg            uint16

	Brand Brand
}
