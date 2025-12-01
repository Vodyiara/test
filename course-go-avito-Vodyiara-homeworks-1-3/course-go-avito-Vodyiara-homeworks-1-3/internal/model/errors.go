package model

import "errors"

var (
	ErrInvalidName   = errors.New("invalid name: name is required")
	ErrInvalidPhone  = errors.New("invalid phone: phone is required")
	ErrInvalidStatus = errors.New("invalid status: must be 'available', 'busy', or 'paused'")
	ErrIDRequired    = errors.New("id is required")
	ErrInvalidID     = errors.New("invalid id format")

	ErrCourierNotFound    = errors.New("courier not found")
	ErrPhoneAlreadyExists = errors.New("phone number already exists")

	ErrInternal = errors.New("internal server error")
)
