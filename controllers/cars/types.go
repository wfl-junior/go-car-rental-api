package controllers

type CarBody struct {
	Brand                 string  `binding:"required"`
	Model                 string  `binding:"required"`
	PriceInUsd            float64 `binding:"required"`
	HorsePower            uint16  `binding:"required"`
	TorqueInLb            uint16  `binding:"required"`
	TopSpeedInKm          uint16  `binding:"required"`
	AccelerationSpeedInKm float32 `binding:"required"`
	WeightInKg            uint16  `binding:"required"`
}