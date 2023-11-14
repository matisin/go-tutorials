package repository

import (
	"errors"
	"users-service/internal/core/dto"
)

var (
	ErrDuplicateIdentification = errors.New("duplicate identification")
	ErrNotFound                = errors.New("user not found")
	ErrDuplicatedEntry         = errors.New("duplicated entry")
	ErrInternalServer          = errors.New("there was a problem with the resquest")
)

type User interface {
	Create(user dto.User) error
	ReadOne(id string) (dto.User, error)
	Read(params map[string]string,
		selects []string,
		orders []string,
		joins []string,
	) ([]dto.User, error)
	// Update(id string, user dto.User) error
	// Delete(id string) error
}
