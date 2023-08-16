package se

import (
	"database/sql"
	"errors"
	"time"
)

type Type int

const (
	Unknown = iota
	ErrConflict
	ErrNotFound
	ErrServer
	ErrBadRequest
	ErrForbidden
)

func (t Type) String() string {
	switch t {
	case ErrConflict:
		return "Conflict"
	case ErrNotFound:
		return "NotFound"
	case ErrServer:
		return "InternalServer"
	case ErrBadRequest:
		return "BadRequest"
	case ErrForbidden:
		return "Forbidden"
	default:
		return "Unknown"

	}
}

type ServiceError struct {
	Time        string `json:"time"`
	Description string `json:"description"`
	Error       any    `json:"error" swaggertype:"string"`
	ErrorType   Type   `json:"type" swaggertype:"integer"`
}

func (se *ServiceError) Type() Type {
	return se.ErrorType
}

func New(description string, err error, errType Type) *ServiceError {
	return &ServiceError{Time: time.Now().Local().Format(time.RFC3339), Description: description, Error: err.Error(), ErrorType: errType}
}

func Internal(err error, descriptions ...string) *ServiceError {
	description := "Internal server error"

	if len(descriptions) > 0 {
		description = descriptions[0]
	}
	return New(description, err, ErrServer)
}

func Validating(err error) *ServiceError {
	return New("Bad input request", err, ErrBadRequest)
}

func Forbidden(description string) *ServiceError {
	return New(description, errors.New("forbidden"), ErrForbidden)
}

func Conflict(description string) *ServiceError {
	return New(description, errors.New("conflict"), ErrConflict)
}

func NotFound(description string) *ServiceError {
	return New(description, errors.New("not found"), ErrNotFound)
}

func BadRequest(description string) *ServiceError {
	return New(description, errors.New("bad request"), ErrBadRequest)
}

func NotFoundOrInternal(err error, descriptions ...string) *ServiceError {
	description := "not found"
	if len(descriptions) > 0 {
		description = descriptions[0]
	}
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return NotFound(description)
	default:
		return Internal(err)
	}
}

func ConflictOrInternal(err error, descriptions ...string) *ServiceError {
	description := "conflict"
	if len(descriptions) > 0 {
		description = descriptions[0]
	}

	switch {
	case errors.Is(err, nil):
		return Conflict(description)
	default:
		return Internal(err)
	}
}

func BadRequestOrInternal(description string, err error) *ServiceError {
	switch {
	case errors.Is(err, nil):
		return BadRequest(description)
	default:
		return Internal(err)
	}
}
