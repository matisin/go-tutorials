package service

import ("users-service/internal/core/dto")

type mockUserRepository struct {
	Data map[string]dto.User
}

func (m *mockUserRepository) Insert(user dto.UserDTO) error {
	// Simulate a duplicate user case
	m.Data[]
	if user.UserName == "test_user" {
		return repository.DuplicateUser
	}

	// Simulate successful insertion
	return nil
}

type mockRepository {

}