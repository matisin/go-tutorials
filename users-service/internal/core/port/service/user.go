package service

import (
	"users-service/internal/core/model/request"
	"users-service/internal/core/model/response"
)

type UserService interface {
	SignUp(request *request.SignUp) *response.Response
	// SignOut(request *request.SignOut) *response.Response
	// Refresh(request *request.Refresh) *response.Response
	// Login(request *request.Login) *response.Response
	// StartContactValidation(request *request.StartContactValidation)
	// ValidateContact(request *request.ValidateContact)
	// CheckUserExists(request *request.CheckUserExists)
	// GetAvailableRoutes(request *request.GetAvailableRoutes)
	// ChangePassword(request *request.ChangePassword)
	// GetUser(request *request.GetUser) *response.Response
}
