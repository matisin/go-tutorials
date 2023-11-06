package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	ID          uuid.UUID    `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Groups      []Group      `gorm:"foreignkey:ProfileID"`
	Permissions []Permission `gorm:"foreignkey:ProfileID"`
	Name        string       `gorm:"type:varchar(50);unique"`
}
