package models

import (
	"time"

	"github.com/google/uuid"
)

type Brand struct {
	Id        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `gorm:"type:varchar(255);uniqueIndex"`

	Cars []Car
}
