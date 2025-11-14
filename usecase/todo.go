package usecase

import (
	"github.com/benyamin218118/todoService/domain"
	"github.com/benyamin218118/todoService/domain/contracts"
)

type TodoUseCase struct {
	repo   contracts.ITodoRepository
	pubsub contracts.IPubSub
}

func NewTodoUseCase(repo contracts.ITodoRepository, pubsub contracts.IPubSub) *TodoUseCase {
	return &TodoUseCase{
		repo:   repo,
		pubsub: pubsub,
	}
}

func (u *TodoUseCase) CreateTodoItem(todo domain.TodoItem) (id *string, err error) {
	if err := todo.Validate(); err != nil {
		return nil, &domain.BadRequestError{
			Msg: err.Error(),
		}
	}

	if id, err := u.repo.Save(todo); err != nil {
		return id, err
	}

	todo.ID = *id

	return id, u.pubsub.Publish("todo_stream", todo)
}
