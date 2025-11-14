package domain

import (
	"errors"
	"time"
)

type TodoItem struct {
	ID          string    `json:"id"` // UUID
	Description string    `json:"description"`
	DueDate     time.Time `json:"dueDate"`
	FileID      *string   `json:"fileId"` // S3 file reference
}

func (t *TodoItem) Validate() error {
	if t.Description == "" {
		return errors.New("description cannot be empty")
	}
	if t.DueDate.Before(time.Now()) {
		return errors.New("due date must be in the future")
	}
	return nil
}
