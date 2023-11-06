package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID       string    `gorm:"type:uuid"`
	Device       string    `gorm:"type:varchar(50)"`
	Active       bool      `gorm:"default:true"`
	IpAddress    string    `gorm:"type:varchar(50)"`
	ProfileID    string    `gorm:"type:uuid"`
	Audits       []Audit   `gorm:"foreignkey:SessionID"`
	SessionStart time.Time
	SessionEnd   time.Time
}
