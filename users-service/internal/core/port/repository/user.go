package repository

import (
	"errors"
	"users-service/internal/core/dto"
)

var (
	ErrDuplicateUser           = errors.New("duplicate user")
	ErrDuplicateMail           = errors.New("duplicate mail")
	ErrDuplicateIdentification = errors.New("duplicate identification")
)

type UserRepository interface {
	Create(user dto.User) error
	// Read(id string) error
	// Update(id string, user dto.User) error
	// Delete(id string) error
}
