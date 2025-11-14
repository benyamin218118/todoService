package contracts

import (
	"context"
	"io"
)

type IStorage interface {
	Upload(file io.Reader, filename string) (*string, error)
	GetFileName(ctx context.Context, id string) (*string, error)
}
