package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type City struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	RegionID string    `gorm:"type:uuid"`
	Name     string    `gorm:"type:varchar(50)"`
}
