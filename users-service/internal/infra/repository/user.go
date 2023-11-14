package repository

import (
	"errors"

	"users-service/internal/core/dto"
	"users-service/internal/core/port/repository"

	"gorm.io/gorm"
)

type UserRepository struct {
	db repository.Database
}

func NewUserRepository(db repository.Database) repository.User {
	return &UserRepository{
		db: db,
	}
}

func (repo UserRepository) Read(
	params map[string]string,
	selects []string,
	orders []string,
	joins []string,
) ([]dto.User, error) {
	db := repo.db.GetDB()

	// Aplicar los parámetros de la consulta
	for key, value := range params {
		db = db.Where(key, value)
	}

	// Aplicar los campos select
	if len(selects) > 0 {
		db = db.Select(selects)
	}

	// Aplicar ordenamiento
	for _, order := range orders {
		db = db.Order(order)
	}

	// Aplicar joins
	for _, join := range joins {
		db = db.Joins(join)
	}

	var users []dto.User
	result := db.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (repo UserRepository) Create(user dto.User) error {
	db := repo.db.GetDB()
	result := db.Create(&user)

	if result.Error != nil {
		return repository.ErrDuplicatedEntry
	}

	return nil
}

func (repo UserRepository) ReadOne(id string) (dto.User, error) {
	var user dto.User
	db := repo.db.GetDB()
	result := db.First(&user, "id = ?", id)

	if result.Error != nil {
		switch {
		case errors.Is(result.Error, gorm.ErrRecordNotFound):
			return dto.User{}, repository.ErrNotFound
		default:
			return dto.User{}, repository.ErrInternalServer
		}
	}
	return user, nil
}
