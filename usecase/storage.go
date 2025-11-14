package usecase

import (
	"io"

	"github.com/benyamin218118/todoService/domain/contracts"
)

type StorageUseCase struct {
	svc contracts.IStorage
}

func NewStorageUseCase(svc contracts.IStorage) *StorageUseCase {
	return &StorageUseCase{
		svc: svc,
	}
}

func (u *StorageUseCase) Upload(file io.Reader, name string) (id *string, err error) {
	return u.svc.Upload(file, name)
}
