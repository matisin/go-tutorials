package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Resource struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID    string    `gorm:"type:uuid"`
	ForeignID string    `gorm:"type:uuid"`
	KeyID     string    `gorm:"type:varchar(50)"`
}
