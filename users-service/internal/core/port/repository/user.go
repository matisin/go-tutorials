package repository

import (
	"errors"
	"users-service/internal/core/dto"
)

var (
	ErrDuplicateUser           = errors.New("duplicate user")
	ErrDuplicateMail           = errors.New("duplicate mail")
	ErrDuplicateIdentification = errors.New("duplicate identification")
	ErrNotFound                = errors.New("user not found")
)

type User interface {
	Create(user dto.User) (string, error)
	ReadOne(id string) (dto.User, error)
	// Update(id string, user dto.User) error
	// Delete(id string) error
}
