package constant

import "errors"

var (
	ErrRecordNotFound      = errors.New("record not found")
	ErrDataAlreadyExist    = errors.New("data already exist")
	ErrStatusInternalError = errors.New("internal server error")
	ErrInvalidToken        = errors.New("token signature is invalid: signature is invalid")
	ErrMissingToken        = errors.New("missing value in request header")
)
