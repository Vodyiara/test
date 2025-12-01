package model

import "time"

type Courier struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Status    string    `json:"status"` // available, busy, paused
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type CreateCourierRequest struct {
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Status string `json:"status"`
}

type UpdateCourierRequest struct {
	ID     *int64  `json:"id"`
	Name   *string `json:"name"`
	Phone  *string `json:"phone"`
	Status *string `json:"status"`
}

func (r *CreateCourierRequest) Validate() error {
	if r.Name == "" {
		return ErrInvalidName
	}
	if r.Phone == "" {
		return ErrInvalidPhone
	}
	if !IsValidStatus(r.Status) {
		return ErrInvalidStatus
	}
	return nil
}

func (r *UpdateCourierRequest) Validate() error {
	if r.ID == nil {
		return ErrIDRequired
	}
	if r.Status != nil && !IsValidStatus(*r.Status) {
		return ErrInvalidStatus
	}
	return nil
}

func IsValidStatus(status string) bool {
	validStatuses := map[string]bool{
		"available": true,
		"busy":      true,
		"paused":    true,
	}
	return validStatuses[status]
}
