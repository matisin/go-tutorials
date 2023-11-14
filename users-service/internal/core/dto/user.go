package dto

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name           string
	Lastname       string
	Mail           string
	State          string `gorm:"type:char(3);default:PEN"`
	Identification string
	Phone          string
	ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
}
