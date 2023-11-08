package users

import (
	"database-api/domains/users/entities"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"database-api/db"
)

func CreateService(user entities.User) (entities.User, error) {

	db := db.GetDBInstance()
	result := db.Create(&user)

	if result.Error != nil {
		return entities.User{}, result.Error
	}

	return user, nil
}

func FindOneService(ID string) (entities.User, error) {
	var user entities.User

	uuid, err := uuid.Parse(ID)
	if err != nil {
		return entities.User{}, ErrInvalidUUID
	}

	db := db.GetDBInstance()
	result := db.First(&user, "id = ?", uuid)

	if result.Error != nil {
		switch {
		case errors.Is(result.Error, gorm.ErrRecordNotFound):
			return entities.User{}, ErrUserNotFound
		default:
			return entities.User{}, ErrInternalServer
		}
	}

	return user, nil
}

func FindAllService() ([]entities.User, error) {
	var users []entities.User
	db := db.GetDBInstance()
	result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func UpdateOneService(ID string, user entities.User) (entities.User, error) {
	db := db.GetDBInstance()

	uuid, err := uuid.Parse(ID)
	if err != nil {
		return entities.User{}, ErrInvalidUUID
	}
	user.ID = uuid

	result := db.Model(&user).Updates(user)

	if result.Error != nil {
		return entities.User{}, ErrInternalServer

	}
	if result.RowsAffected == 0 {
		return entities.User{}, ErrUserNotFound
	}

	return user, nil
}

func DeleteOneService(ID string) (entities.User, error) {
	uuid, err := uuid.Parse(ID)
	if err != nil {
		return entities.User{}, ErrInvalidUUID
	}

	user := entities.User{ID: uuid}

	db := db.GetDBInstance()
	result := db.Delete(&user)

	if result.Error != nil {
		return entities.User{}, ErrInternalServer

	}
	if result.RowsAffected == 0 {
		return entities.User{}, ErrUserNotFound
	}

	return user, nil
}
