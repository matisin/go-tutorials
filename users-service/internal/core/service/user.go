package service

import (
	"log"
	"users-service/internal/core/dto"
	"users-service/internal/core/entity/error_code"
	"users-service/internal/core/model/request"
	"users-service/internal/core/model/response"
	"users-service/internal/core/port/repository"
	"users-service/internal/core/port/service"
)

const (
	invalidUserNameErrMsg = "invalid username"
	invalidPasswordErrMsg = "invalid password"
	noParamsErrMsg        = "there's no params for query"
)

type userService struct {
	userRepo repository.User
}

func NewUserService(userRepo repository.User) service.UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u userService) CheckUserExists(request *request.CheckUserExists) *response.Response {
	var params map[string]string
	selects := []string{}
	joins := []string{}
	orders := []string{"name"}

	switch {
	case request.Identification != "":
		params = map[string]string{
			"identification = ?": request.Identification,
		}
	case request.Mail != "":
		params = map[string]string{
			"mail = ?": request.Identification,
		}
	case request.Phone != "":
		params = map[string]string{
			"phone = ?": request.Identification,
		}
	default:
		return u.createFailedResponse(error_code.InvalidRequest, noParamsErrMsg)
	}
	users, _ := u.userRepo.Read(params, selects, orders, joins)

	log.Println(users)

	CheckUserExistData := response.CheckUserExistData{
		Exists: len(users) == 0,
	}
	return u.createSuccessResponse(CheckUserExistData)
}

func (u userService) GetUser(request *request.GetUser) *response.Response {

	user, _ := u.userRepo.ReadOne(request.UserId)

	// create data response
	GetUserData := response.GetUserData{
		Name:           user.Name,
		Lastname:       user.Lastname,
		Mail:           user.Mail,
		State:          user.State,
		Identification: user.Identification,
		Phone:          user.Phone,
		ID:             user.ID.String(),
	}

	return u.createSuccessResponse(GetUserData)
}

func (u userService) SignUp(request *request.SignUp) *response.Response {
	// validate request
	if len(request.Mail) == 0 {
		return u.createFailedResponse(error_code.InvalidRequest, invalidUserNameErrMsg)
	}

	// if len(request.P) == 0 {
	// 	return u.createFailedResponse(error_code.InvalidRequest, invalidPasswordErrMsg)
	// }

	// currentTime := utils.GetUTCCurrentMillis()
	userDTO := dto.User{
		Name:           request.Name,
		Lastname:       request.Lastname,
		Mail:           request.Mail,
		Identification: request.Identification,
		Phone:          request.Phone,
		// CreatedAt:   currentTime,
		// UpdatedAt:   currentTime,
	}

	// save a new user
	err := u.userRepo.Create(userDTO)
	if err != nil {
		if err == repository.ErrDuplicatedEntry {
			return u.createFailedResponse(error_code.MailAlreadyExists, err.Error())
		}

		return u.createFailedResponse(error_code.InternalError, error_code.InternalErrMsg)
	}

	// create data response
	signUpData := response.SignUpData{
		Name: userDTO.Name,
	}

	return u.createSuccessResponse(signUpData)
}

func (u userService) createFailedResponse(
	code error_code.ErrorCode, message string,
) *response.Response {
	return &response.Response{
		ErrorCode:    code,
		ErrorMessage: message,
		Status:       false,
	}
}

func (u userService) createSuccessResponse(data interface{}) *response.Response {
	return &response.Response{
		Data:         data,
		ErrorCode:    error_code.Success,
		ErrorMessage: error_code.SuccessErrMsg,
		Status:       true,
	}
}
