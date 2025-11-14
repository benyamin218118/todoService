package usecase_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/benyamin218118/todoService/domain"
	"github.com/benyamin218118/todoService/infra/repositories/mocks"
	"github.com/benyamin218118/todoService/usecase"
	"github.com/golang/mock/gomock"
)

func TestCreateTodoItem_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mocks.NewMockITodoRepository(ctrl)
	pubsubMock := mocks.NewMockIPubSub(ctrl)
	storageMock := mocks.NewMockIStorage(ctrl)

	todoUseCase := usecase.NewTodoUseCase(repoMock, storageMock, pubsubMock)

	todo := domain.TodoItem{
		ID:          "uuid-123",
		Description: "Test Todo",
		DueDate:     time.Now().Add(time.Hour),
		FileID:      nil,
	}
	data, err := StructToMap(todo)
	if err != nil {
		t.Fatal("structToMap failed", err)
	}
	repoMock.EXPECT().Save(gomock.Any(), todo).Return(&todo.ID, nil)
	pubsubMock.EXPECT().Publish("todoCreated", data).Return(nil)

	// Call the use case
	createdTodoItem, err := todoUseCase.CreateTodoItem(context.Background(), todo)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if createdTodoItem == nil {
		t.Fatalf("created todo item should not be nil")
	}
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

func TestCreateTodoItem_FailSave(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mocks.NewMockITodoRepository(ctrl)
	redisMock := mocks.NewMockIPubSub(ctrl)
	storageMock := mocks.NewMockIStorage(ctrl)

	todoUseCase := usecase.NewTodoUseCase(repoMock, storageMock, redisMock)

	todo := domain.TodoItem{
		Description: "",
		DueDate:     time.Now().Add(time.Hour),
	}

	// repoMock.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil, domain.BadRequestError{Msg: "description cannot be empty"})

	_, err := todoUseCase.CreateTodoItem(context.Background(), todo)
	if err == nil {
		t.Fatal("expected error but got nil")
	}
}
