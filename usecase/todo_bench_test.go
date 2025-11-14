package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/benyamin218118/todoService/domain"
	"github.com/benyamin218118/todoService/infra/repositories/mocks"
	"github.com/benyamin218118/todoService/usecase"
	"github.com/golang/mock/gomock"
)

func BenchmarkCreateTodoItem(b *testing.B) {
	ctrl := gomock.NewController(b)
	defer ctrl.Finish()

	repoMock := mocks.NewMockITodoRepository(ctrl)
	redisMock := mocks.NewMockIPubSub(ctrl)
	storage := mocks.NewMockIStorage(ctrl)

	todoUseCase := usecase.NewTodoUseCase(repoMock, storage, redisMock)

	todo := domain.TodoItem{
		ID:          "uuid-123",
		Description: "Benchmark Todo",
		DueDate:     time.Now().Add(time.Hour),
		FileID:      nil,
	}

	// Tell mocks to allow any number of calls
	repoMock.EXPECT().Save(gomock.Any(), todo).AnyTimes().Return(&todo.ID, nil)
	data, err := StructToMap(todo)
	if err != nil {
		b.Fatal(err)
	}
	redisMock.EXPECT().Publish("todoCreated", data).AnyTimes().Return(nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		todoUseCase.CreateTodoItem(context.Background(), todo)
	}
}
