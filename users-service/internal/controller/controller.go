package controller

import (
	"net/http"

	"users-service/internal/core/entity/error_code"
	"users-service/internal/core/model/request"
	"users-service/internal/core/model/response"
	"users-service/internal/core/port/service"

	"github.com/gin-gonic/gin"
)

var (
	invalidRequestResponse = &response.Response{
		ErrorCode:    error_code.InvalidRequest,
		ErrorMessage: error_code.InvalidRequestErrMsg,
		Status:       false,
	}
)

type UserController struct {
	gin         *gin.Engine
	userService service.UserService
}

func NewUserController(
	gin *gin.Engine,
	userService service.UserService,
) UserController {
	return UserController{
		gin:         gin,
		userService: userService,
	}

}

func (u UserController) InitRouter() {
	api := u.gin.Group("/api/v1/users")
	api.GET("/id/:id", u.getUser)
	api.POST("/signup", u.signUp)
	api.GET("/check_users", u.checkUsers)
}

func (u UserController) signUp(c *gin.Context) {
	req, err := u.parseRequest(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &invalidRequestResponse)
		return
	}

	resp := u.userService.SignUp(req)
	c.JSON(http.StatusOK, resp)
}

func (u UserController) parseRequest(ctx *gin.Context) (*request.SignUp, error) {
	var req request.SignUp
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	return &req, nil
}

func (u UserController) getUser(c *gin.Context) {
	id := c.Param("id")
	req := request.GetUser{UserId: id}
	// if err := c.ShouldBindJSON(&req); err != nil {
	// 	c.AbortWithStatusJSON(http.StatusOK, &invalidRequestResponse)
	// 	return
	// }

	resp := u.userService.GetUser(&req)
	c.JSON(http.StatusOK, resp)
}

func (u UserController) checkUsers(c *gin.Context) {
	var req request.CheckUserExists
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &invalidRequestResponse)
		return
	}
	// if err := c.ShouldBindJSON(&req); err != nil {
	// 	c.AbortWithStatusJSON(http.StatusOK, &invalidRequestResponse)
	// 	return
	// }

	resp := u.userService.CheckUserExists(&req)
	c.JSON(http.StatusOK, resp)
}
