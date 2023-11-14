package response

import "users-service/internal/core/entity/error_code"

type Response struct {
	Data         interface{}          `json:"data"`
	Status       bool                 `json:"status"`
	ErrorCode    error_code.ErrorCode `json:"errorCode"`
	ErrorMessage string               `json:"errorMessage"`
}

type SignUpData struct {
	Name string `json:"displayName"`
}

type GetUserData struct {
	Name           string `json:"name"`
	Lastname       string `json:"lastname"`
	Mail           string `json:"mail"`
	State          string `json:"state"`
	Identification string `json:"identification"`
	Phone          string `json:"phone"`
	ID             string `json:"id"`
}

type CheckUserExistData struct {
	Exists bool `json:"exists"`
}
