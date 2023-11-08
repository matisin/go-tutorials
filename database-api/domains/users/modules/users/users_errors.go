package users

import "errors"

var ErrInvalidUUID = errors.New("invalid UUID ID")

var ErrUserNotFound = errors.New("user not found")

var ErrInternalServer = errors.New("there was a problem with the request")

var ErrUserDataNotValid = errors.New("user data not valid")
