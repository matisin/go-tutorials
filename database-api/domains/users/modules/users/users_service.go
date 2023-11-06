package users

import (
	"database-api/domains/users/entities"

	"database-api/db"
)

func CreateService(data entities.User) (entities.User, error) {

	db := db.GetDBInstance()
	result := db.Create(&data)

	if result.Error != nil {
		return entities.User{}, result.Error
	}

	return data, nil
}

func FindOneService(ID string) (entities.User, error) {
	var user entities.User

	db := db.GetDBInstance()
	result := db.First(&user, "id = ?", ID)

	if result.Error != nil {
		return entities.User{}, result.Error
	}

	return user, nil
}
