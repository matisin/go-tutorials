package users

import (
	"database-api/domains/users/entities"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateController(ctx *gin.Context) {
	var user entities.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": ErrUserDataNotValid.Error()})
		return
	}

	userCreated, err := CreateService(user)
	if err != nil {
		switch {
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServer.Error()})
		}
		return
	}

	ctx.JSON(http.StatusCreated, userCreated)
}

func FindOneController(ctx *gin.Context) {
	userID := ctx.Param("id")

	fetchedUser, err := FindOneService(userID)
	if err != nil {
		switch {
		case errors.Is(err, ErrInvalidUUID):
			ctx.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidUUID.Error()})
		case errors.Is(err, ErrUserNotFound):
			ctx.JSON(http.StatusNotFound, gin.H{"error": ErrUserNotFound.Error()})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServer.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, fetchedUser)
}

func FindAllController(ctx *gin.Context) {
	fetchedUsers, err := FindAllService()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServer.Error()})
		return
	}

	ctx.JSON(http.StatusOK, fetchedUsers)
}

func UpdateOneController(ctx *gin.Context) {
	ID := ctx.Param("id")

	var user entities.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": ErrUserDataNotValid.Error()})
		return
	}
	updatedUser, err := UpdateOneService(ID, user)

	if err != nil {
		switch {
		case errors.Is(err, ErrInvalidUUID):
			ctx.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidUUID.Error()})
		case errors.Is(err, ErrUserNotFound):
			ctx.JSON(http.StatusNotFound, gin.H{"error": ErrUserNotFound.Error()})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServer.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, updatedUser)
}

func DeleteOneController(ctx *gin.Context) {
	id := ctx.Param("id")
	deletedUser, err := DeleteOneService(id)
	if err != nil {
		switch {
		case errors.Is(err, ErrInvalidUUID):
			ctx.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidUUID.Error()})
		case errors.Is(err, ErrUserNotFound):
			ctx.JSON(http.StatusNotFound, gin.H{"error": ErrUserNotFound.Error()})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServer.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, deletedUser)
}
