package usecase_test

import (
	"bytes"
	"testing"

	"github.com/benyamin218118/todoService/infra/repositories/mocks"
	"github.com/benyamin218118/todoService/usecase"
	"github.com/golang/mock/gomock"
)

func BenchmarkUploadFile(b *testing.B) {
	ctrl := gomock.NewController(b)
	defer ctrl.Finish()

	s3Mock := mocks.NewMockIStorage(ctrl)

	storageUseCase := usecase.NewStorageUseCase(s3Mock)

	fileContent := bytes.NewReader([]byte("benchmark content"))
	filename := "benchmark.txt"

	s3Mock.EXPECT().Upload(fileContent, filename).AnyTimes().DoAndReturn(func(any, any) (any, any) {
		return filename, nil
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = storageUseCase.Upload(fileContent, filename)
	}
}
