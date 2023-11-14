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
		if element.Mail == user.Mail || element.Identification == user.Identification {
			return repository.ErrDuplicatedEntry
		}
	}

	uuid, _ := uuid.NewRandom()
	strUUID := uuid.String()
	user.ID = uuid
	m.Data[strUUID] = user
	return nil
}

func (m *mockUserRepository) ReadOne(id string) (dto.User, error) {
	if m.Data == nil {
		m.Data = make(map[string]dto.User)
	}
	user, ok := m.Data[id]
	if ok {
		return user, nil
	} else {
		return dto.User{}, repository.ErrNotFound
	}
}

func (m *mockUserRepository) RestartData() {
	m.Data = make(map[string]dto.User)
}
