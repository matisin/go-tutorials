package repository

import (
	"errors"
	"users-service/internal/core/dto"
)

var (
	ErrDuplicateUser = errors.New("duplicate user")
)

type UserRepository interface {
	Create(user dto.User) error
	Read(id string) error
	Update(id string, user dto.User) error
	Delete(id string) error
}
