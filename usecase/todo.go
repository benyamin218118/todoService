package usecase

import (
	"context"
	"encoding/json"

	"github.com/benyamin218118/todoService/domain"
	"github.com/benyamin218118/todoService/domain/contracts"
)

type TodoUseCase struct {
	todoRepo    contracts.ITodoRepository
	storageRepo contracts.IStorage
	pubsub      contracts.IPubSub
}

func NewTodoUseCase(todoRepo contracts.ITodoRepository, storageRepo contracts.IStorage, pubsub contracts.IPubSub) *TodoUseCase {
	return &TodoUseCase{
		todoRepo:    todoRepo,
		storageRepo: storageRepo,
		pubsub:      pubsub,
	}
}

func (u *TodoUseCase) CreateTodoItem(ctx context.Context, todo domain.TodoItem) (id *string, err error) {
	if err := todo.Validate(); err != nil {
		return nil, &domain.BadRequestError{
			Msg: err.Error(),
		}
	}

	if todo.FileID != nil && len(*todo.FileID) > 0 {
		_, err := u.storageRepo.GetFileName(ctx, *todo.FileID)
		if err != nil {
			return nil, domain.ForbiddenError{Msg: `couldnt find any file with this id`}
		}
	}

	if id, err = u.todoRepo.Save(ctx, todo); err != nil {
		return id, err
	}

	todo.ID = *id

	data, err := StructToMap(todo)
	if err != nil {
		return id, err
	}
	return id, u.pubsub.Publish("todoCreated", data)
}

func StructToMap(s any) (map[string]interface{}, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	err = json.Unmarshal(b, &m)
	return m, err
}
