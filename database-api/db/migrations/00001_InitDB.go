package migration

import (
	"database-api/db"
	userEntities "database-api/domains/users/entities"
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up00001, Down00001)
}

func Up00001(tx *sql.Tx) error {
	err := db.GetDBInstance().AutoMigrate(
		&userEntities.Module{},
		&userEntities.Action{},
		&userEntities.Profile{},
		&userEntities.User{},
		&userEntities.Session{},
		&userEntities.Audit{},
		&userEntities.Group{},
		&userEntities.Permission{},
		&userEntities.Resource{},
	)
	if err != nil {
		return err
	}
	return nil
}

func Down00001(tx *sql.Tx) error {
	// Aquí puedes revertir los cambios realizados en Up00001.
	return nil
}
