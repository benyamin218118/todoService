package domain

import (
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
		return BadRequestError{Msg: "description cannot be empty"}
	}
	if len(t.Description) > 10240 {
		return BadRequestError{Msg: "description len should be < 10240"}
	}
	if t.DueDate.Before(time.Now()) {
		return BadRequestError{Msg: "due date must be in the future"}
	}
	return nil
}
