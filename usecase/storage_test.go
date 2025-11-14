package usecase_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/benyamin218118/todoService/infra/repositories/mocks"
	"github.com/benyamin218118/todoService/usecase"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

func TestUploadFile_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	s3Mock := mocks.NewMockIStorage(ctrl) // Storage interface

	fileContent := bytes.NewReader([]byte("hello world"))
	filename := "test.txt"

	s3Mock.EXPECT().Upload(fileContent, filename).DoAndReturn(func(io.Reader, string) (*string, error) {
		fileId := uuid.NewString()
		return &fileId, nil
	})

	storageUseCase := usecase.NewStorageUseCase(s3Mock)

	fileID, err := storageUseCase.Upload(fileContent, filename)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if uuid.Validate(*fileID) != nil {
		t.Fatalf("unexpected fileID: %v", fileID)
	}
}

func TestUploadFile_Fail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	s3Mock := mocks.NewMockIStorage(ctrl) // Storage interface

	var fileContent io.Reader
	filename := "test.txt"

	storageUseCase := usecase.NewStorageUseCase(s3Mock)

	_, err := storageUseCase.Upload(fileContent, filename)
	if err == nil {
		t.Fatal("expected error but got nil")
	}
}
