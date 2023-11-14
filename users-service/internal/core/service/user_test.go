package service

import (
	"testing"

	"users-service/internal/core/entity/error_code"
	"users-service/internal/core/model/request"
	"users-service/internal/core/model/response"
)

var (
	UserRepo    = &mockUserRepository{}
	UserService = NewUserService(UserRepo)
)

func TestUserService_GetUser(t *testing.T) {
	UserRepo.RestartData()

	UserService.SignUp(&request.SignUp{
		Name:           "Pedro",
		Lastname:       "Perez",
		Mail:           "pedro_perez@gmail.com",
		Identification: "182596086",
		Phone:          "+56987334383",
	})

	UserService.SignUp(&request.SignUp{
		Name:           "Juan",
		Lastname:       "Perez",
		Mail:           "juan_perez@gmail.com",
		Identification: "182596087",
		Phone:          "+56987334383",
	})

	var uuid string
	for key := range UserRepo.Data {
		uuid = key
		break
	}

	req := &request.GetUser{
		UserId: uuid,
	}

	res := UserService.GetUser(req)
	if !res.Status {
		t.Errorf("expected status to be true, got false")
	}

	data := res.Data.(response.GetUserData)
	if data.ID == "" {
		t.Errorf("expected non-empty display name, got empty")
	}

}

func TestUserService_SignUp_Success(t *testing.T) {
	// Restart DB
	UserRepo.RestartData()

	// Test case: Successful signup
	req := &request.SignUp{
		Name:           "Pedro",
		Lastname:       "Perez",
		Mail:           "pedro_perez@gmail.com",
		Identification: "182596086",
		Phone:          "+56987334383",
	}

	res := UserService.SignUp(req)
	if !res.Status {
		t.Errorf("expected status to be true, got false")
	}

	data := res.Data.(response.SignUpData)
	if data.Name == "" {
		t.Errorf("expected non-empty display name, got empty")
	}
}

func TestUserService_SignUp_DuplicateUser(t *testing.T) {
	// Create a mock User for testing
	UserRepo.RestartData()
	// Test case: Successful signup
	req := &request.SignUp{
		Name:           "Pedro",
		Lastname:       "Perez",
		Mail:           "pedro_perez@gmail.com",
		Identification: "182596086",
		Phone:          "+56987334383",
	}

	res := UserService.SignUp(req)
	if !res.Status {
		t.Errorf("expected status to be true, got false")
	}

	req = &request.SignUp{
		Name:           "Pedro",
		Lastname:       "Perez",
		Mail:           "pedro_perez@gmail.com",
		Identification: "182596086",
		Phone:          "+56987334383",
	}

	res = UserService.SignUp(req)

	if res.ErrorCode != error_code.MailAlreadyExists {
		t.Errorf("expected error code to be MailAlreadyExists, got %s", res.ErrorCode)
	}

}

// func TestUserService_SignUp_InvalidUsername(t *testing.T) {
// 	// Create a mock User for testing
// 	userRepo := &mockUserRepository{}

// 	// Create the UserService using the mock User
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
// 	// Create a mock User for testing
// 	userRepo := &mockUserRepository{}

// 	// Create the UserService using the mock User
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
// 	// Create a mock User for testing
// 	userRepo := &mockUserRepository{}

// 	// Create the UserService using the mock User
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
