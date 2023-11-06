package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Audit struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	SessionID  string    `gorm:"type:uuid"`
	ActionID   string    `gorm:"type:uuid"`
	ResourceID string    `gorm:"type:uuid"`
	Input      string    `gorm:"type:jsonb"`
	Output     string    `gorm:"type:jsonb"`
}
