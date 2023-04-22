package models

import (
	"time"

	"github.com/google/uuid"
)

type Car struct {
	Id       							uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
	Brand                 string
	Model 								string
	PriceInUsd            float64
	HorsePower            uint16
	TorqueInLb            uint16
	TopSpeedInKm          uint16
	AccelerationSpeedInKm float32
	WeightInKg            uint16
}