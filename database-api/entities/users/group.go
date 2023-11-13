package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID    string    `gorm:"type:uuid"`
	ProfileID string    `gorm:"type:uuid"`
}
