package dto

type UserDTO struct {
	UserName    string
	Password    string
	DisplayName string
	CreatedAt   uint64
	UpdatedAt   uint64
}

type User struct {
	Name           string `gorm:"type:varchar(100)"`
	Lastname       string `gorm:"type:varchar(100)"`
	Mail           string `gorm:"type:varchar(100);unique"`
	State          string `gorm:"type:char(3)"`
	Identification string `gorm:"type:varchar(50);unique"`
	Phone          string `gorm:"type:varchar(16)"`
	// Sessions       []Session  `gorm:"foreignkey:UserID"`
	// Groups         []Group    `gorm:"foreignkey:UserID"`
	// Resources      []Resource `gorm:"foreignkey:UserID"`
}
