package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Module struct {
	gorm.Model
	ID      uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Actions []Action  `gorm:"foreignkey:ModuleID"`
	Name    string    `gorm:"type:varchar(50);unique"`
	Url     *string   `gorm:"type:varchar(50)"`
	Path    *string   `gorm:"type:varchar(50)"`
}
