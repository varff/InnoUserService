package apperrors

import (
	"errors"
	"fmt"
)

var (
	ErrWrongPass     = errors.New("wrong password")
	ErrPhoneNotFound = errors.New("invalid phone number")
	ErrPhoneTaken    = errors.New("phone is already taken")
	ErrSingingMethod = errors.New("signing method error")
	ErrWrongClaims   = errors.New("wrong claims")
	ErrUpdate        = errors.New("update error")
	ErrInvalidToken  = errors.New("invalid token")
	ErrTokenSigned   = errors.New("token creation error")
	ErrTokenParsing  = errors.New("token parsing error")
)

func Wrapper(message error, err error) error {
	wrap := fmt.Errorf("%w: %s", message, err)
	return wrap
}

func UnWrapper(err error) error {
	e := errors.Unwrap(err)
	if e != nil {
		return e
	}
	return err
}
