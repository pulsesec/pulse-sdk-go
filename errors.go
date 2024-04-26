package pulse

import (
	"errors"
	"fmt"
)

var (
	ErrTokenNotFound = errors.New("user token not found")
	ErrTokenUsed     = errors.New("user token already used")
	ErrOther         = errors.New("user classification failed")

	errorCodeMap = map[string]error{
		"TOKEN_NOT_FOUND": ErrTokenNotFound,
		"TOKEN_USED":      ErrTokenUsed,
	}
)

type Error struct {
	Message string `json:"error"`
	Code    string `json:"code"`
}

func (e Error) Error() string {
	return e.Message
}

func (e Error) Unwrap() error {
	err, ok := errorCodeMap[e.Code]
	if !ok {
		return ErrOther
	}
	return err
}

// type assertion
var _ error = &Error{}

type errorResponse struct {
	Errors []Error `json:"errors"`
}

func (e errorResponse) Error() error {
	if len(e.Errors) == 0 {
		return nil
	}

	err := error(e.Errors[0])
	for _, e := range e.Errors[1:] {
		err = fmt.Errorf("%w; %w", err, error(e))
	}

	return err
}
