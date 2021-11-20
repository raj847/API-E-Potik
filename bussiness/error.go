package bussiness

import "errors"

var (
	ErrDuplicateData    = errors.New("Data already exist")
	ErrInvalidLoginInfo = errors.New("Username or password is invalid")
	ErrInternalServer   = errors.New("Something went wrong")
	ErrNotFound         = errors.New("Not found")
	ErrUnauthorized     = errors.New("User Unauthorized")
)
