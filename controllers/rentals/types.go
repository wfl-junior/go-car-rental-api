package controllers

import (
	"time"

	"github.com/google/uuid"
)

type RentalBody struct {
	CarId    uuid.UUID `binding:"required"`
	StartsAt time.Time `binding:"required"`
	EndsAt   time.Time `binding:"required"`
}
