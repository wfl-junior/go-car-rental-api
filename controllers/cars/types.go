package controllers

import "github.com/google/uuid"

type CarBody struct {
	BrandId               uuid.UUID `binding:"required"`
	Model                 string    `binding:"required"`
	RentalPriceDailyInUsd float64   `binding:"required"`
	HorsePower            uint16    `binding:"required"`
	TorqueInLb            float32   `binding:"required"`
	TopSpeedInKm          uint16    `binding:"required"`
	AccelerationSpeedInKm float32   `binding:"required"`
	WeightInKg            uint16    `binding:"required"`
}
