package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Country struct {
	gorm.Model
	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Regions      []Region  `gorm:"foreignkey:CountryID"`
	Name         string    `gorm:"type:varchar(50);unique"`
	Abbreviation string    `gorm:"type:char(2);unique"`
}
