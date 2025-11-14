package usecase

import (
	"io"

	"github.com/benyamin218118/todoService/domain"
	"github.com/benyamin218118/todoService/domain/contracts"
)

type StorageUseCase struct {
	repo contracts.IStorage
}

func NewStorageUseCase(repo contracts.IStorage) *StorageUseCase {
	return &StorageUseCase{
		repo: repo,
	}
}

func (u *StorageUseCase) Upload(file io.Reader, name string) (id *string, err error) {
	if file == nil || name == "" {
		return nil, domain.BadRequestError{
			Msg: "invalid file/name",
		}
	}
	return u.repo.Upload(file, name)
}
