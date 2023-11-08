package entities

import (
	"log"

	"gorm.io/gorm"
)

func AutoMigrateEntities(db *gorm.DB) error {
	log.Println("Starting automigration of entities of users domain")

	err := db.AutoMigrate(
		&Country{},
		&Region{},
		&City{},
	)
	if err != nil {
		return err
	}
	log.Println("Succesful automigration of users domain")
	return nil
}
