package contracts

import "github.com/benyamin218118/todoService/domain"

type ITodoRepository interface {
	Save(todo domain.TodoItem) (id *string, err error)
	GetByID(id string) (*domain.TodoItem, error)
}
