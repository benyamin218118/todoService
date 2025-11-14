package repositories_test

import (
	"encoding/json"
	"testing"

	"github.com/benyamin218118/todoService/domain"
	"github.com/benyamin218118/todoService/infra/repositories/mocks"
	"github.com/golang/mock/gomock"
)

func BenchmarkRedisPublish(b *testing.B) {
	ctrl := gomock.NewController(b)
	defer ctrl.Finish()

	redisMock := mocks.NewMockIPubSub(ctrl)

	fid := "fid123"
	todo := domain.TodoItem{
		ID:          "uuid-123",
		Description: "Redis Benchmark",
		FileID:      &fid,
	}

	data, err := StructToMap(todo)
	if err != nil {
		b.Fatal(err)
	}
	redisMock.EXPECT().Publish("todoCreated", data).AnyTimes().Return(nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = redisMock.Publish("todoCreated", data)
	}
}

func StructToMap(s any) (map[string]any, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	var m map[string]any
	err = json.Unmarshal(b, &m)
	return m, err
}
