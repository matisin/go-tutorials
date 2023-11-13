package service

import (
	"users-service/internal/core/dto"
	"users-service/internal/core/port/repository"

	"github.com/google/uuid"
)

type mockUserRepository struct {
	Data map[string]dto.User
}

func (m *mockUserRepository) Create(user dto.User) error {
	// Simulate a duplicate user case
	if m.Data == nil {
		m.Data = make(map[string]dto.User)
	}
	for _, element := range m.Data {
		if element.Mail == user.Mail {
			return repository.ErrDuplicateMail
		}
		if element.Identification == user.Identification {
			return repository.ErrDuplicateIdentification
		}
	}
	stringUuid := uuid.New().String()
	m.Data[stringUuid] = user
	return nil
}
