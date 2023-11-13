package service

import (
	"log"
	"users-service/internal/core/dto"
	"users-service/internal/core/port/repository"

	"github.com/google/uuid"
)

type mockUserRepository struct {
	Data map[string]dto.User
}

func (m *mockUserRepository) Create(user dto.User) (string, error) {
	// Simulate a duplicate user case
	if m.Data == nil {
		m.Data = make(map[string]dto.User)
	}
	for _, element := range m.Data {
		if element.Mail == user.Mail {
			return "", repository.ErrDuplicateMail
		}
		if element.Identification == user.Identification {
			return "", repository.ErrDuplicateIdentification
		}
	}

	uuid, _ := uuid.NewRandom()
	strUUID := uuid.String()
	user.ID = uuid
	m.Data[strUUID] = user
	log.Println(strUUID)
	return strUUID, nil
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
