package usecase

import (
	"io"

	"github.com/benyamin218118/todoService/domain"
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
	if file == nil || name == "" {
		return nil, domain.BadRequestError{
			Msg: "invalid file/name",
		}
	}
	return u.svc.Upload(file, name)
}
