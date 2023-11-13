package service

import (
	"log"
	"testing"

	"users-service/internal/core/entity/error_code"
	"users-service/internal/core/model/request"
	"users-service/internal/core/model/response"
)

// Define a mock UserRepository for testing
// type mockUserRepository struct{}

// func (m *mockUserRepository) Insert(user dto.UserDTO) error {
// 	// Simulate a duplicate user case
// 	if user.UserName == "test_user" {
// 		return repository.DuplicateUser
// 	}

// 	// Simulate successful insertion
// 	return nil
// }

func TestUserService_SignUp_Success(t *testing.T) {
	// Create a mock UserRepository for testing
	userRepo := &mockUserRepository{}

	// Create the UserService using the mock UserRepository
	userService := NewUserService(userRepo)

	// Test case: Successful signup
	req := &request.SignUp{
		Name:           "Pedro",
		Lastname:       "Perez",
		Mail:           "pedro_perez@gmail.com",
		State:          "APR",
		Identification: "182596086",
		Phone:          "+56987334383",
	}

	res := userService.SignUp(req)
	log.Println(res)
	if !res.Status {
		t.Errorf("expected status to be true, got false")
	}

	data := res.Data.(response.SignUpData)
	if data.Name == "" {
		t.Errorf("expected non-empty display name, got empty")
	}
}

func TestUserService_SignUp_DuplicateUser(t *testing.T) {
	// Create a mock UserRepository for testing
	userRepo := &mockUserRepository{}

	// Create the UserService using the mock UserRepository
	userService := NewUserService(userRepo)

	// Test case: Successful signup
	req := &request.SignUp{
		Name:           "Pedro",
		Lastname:       "Perez",
		Mail:           "pedro_perez@gmail.com",
		State:          "APR",
		Identification: "182596086",
		Phone:          "+56987334383",
	}

	res := userService.SignUp(req)
	if !res.Status {
		t.Errorf("expected status to be true, got false")
	}

	req = &request.SignUp{
		Name:           "Pedro",
		Lastname:       "Perez",
		Mail:           "pedro_perez@gmail.com",
		State:          "APR",
		Identification: "182596086",
		Phone:          "+56987334383",
	}

	res = userService.SignUp(req)

	if res.ErrorCode != error_code.MailAlreadyExists {
		t.Errorf("expected error code to be MailAlreadyExists, got %s", res.ErrorCode)
	}
}

// func TestUserService_SignUp_InvalidUsername(t *testing.T) {
// 	// Create a mock UserRepository for testing
// 	userRepo := &mockUserRepository{}

// 	// Create the UserService using the mock UserRepository
// 	userService := NewUserService(userRepo)

// 	// Test case: Invalid request with empty username
// 	req := &request.SignUp{
// 		Name:           "",
// 		Lastname:       "Perez",
// 		Mail:           "pedro_perez@gmail.com",
// 		State:          "APR",
// 		Identification: "182596086",
// 		Phone:          "+56987334383",
// 	}

// 	res := userService.SignUp(req)
// 	if res.Status {
// 		t.Errorf("expected status to be false, got true")
// 	}

// 	if res.ErrorCode != error_code.InvalidRequest {
// 		t.Errorf("expected error code to be InvalidRequest, got %s", res.ErrorCode)
// 	}
// }

// func TestUserService_SignUp_InvalidPassword(t *testing.T) {
// 	// Create a mock UserRepository for testing
// 	userRepo := &mockUserRepository{}

// 	// Create the UserService using the mock UserRepository
// 	userService := NewUserService(userRepo)

// 	// Test case: Invalid request with empty password
// 	req := &request.SignUpRequest{
// 		Username: "test_user",
// 		Password: "",
// 	}

// 	res := userService.SignUp(req)
// 	if res.Status {
// 		t.Errorf("expected status to be false, got true")
// 	}
// 	if res.ErrorCode != error_code.InvalidRequest {
// 		t.Errorf("expected error code to be InvalidRequest, got %s", res.ErrorCode)
// 	}
// }

// func TestUserService_SignUp_DuplicateUser(t *testing.T) {
// 	// Create a mock UserRepository for testing
// 	userRepo := &mockUserRepository{}

// 	// Create the UserService using the mock UserRepository
// 	userService := NewUserService(userRepo)

// 	// Test case: Duplicate user
// 	req := &request.SignUpRequest{
// 		Username: "test_user",
// 		Password: "12345",
// 	}

// 	res := userService.SignUp(req)
// 	if res.Status {
// 		t.Errorf("expected status to be false, got true")
// 	}

// 	if res.ErrorCode != error_code.DuplicateUser {
// 		t.Errorf("expected error code to be DuplicateUser, got %s", res.ErrorCode)
// 	}
// }
