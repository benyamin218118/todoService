package contracts

import (
	"context"

	"github.com/benyamin218118/todoService/domain"
)

type ITodoRepository interface {
	Save(ctx context.Context, todo domain.TodoItem) (id *string, err error)
}
