package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Permission struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	ActionID  string    `gorm:"type:uuid"`
	ProfileID string    `gorm:"type:uuid"`
	Enabled   bool      `gorm:"default:true"`
}
