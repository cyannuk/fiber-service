package errors

import (
	"errors"
)

var (
	ErrNoParameter = errors.New("request parameter not found")
	ErrInvalidParameter = errors.New("invalid request parameter")
)
