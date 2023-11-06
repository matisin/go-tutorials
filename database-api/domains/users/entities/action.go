package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Action struct {
	gorm.Model
	ID          uuid.UUID    `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Permissions []Permission `gorm:"foreignkey:ActionID"`
	Audits      []Audit      `gorm:"foreignkey:ActionID"`
	ModuleID    string       `gorm:"type:uuid"`
	Name        string       `gorm:"type:varchar(50)"`
}
