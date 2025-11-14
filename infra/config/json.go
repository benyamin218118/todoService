package config

import (
	"errors"

	"github.com/benyamin218118/todoService/domain"
	"github.com/benyamin218118/todoService/domain/contracts"
)

type jsonConfigReader struct {
}

func getJSONConfigReader() contracts.IConfigReader {
	return &jsonConfigReader{}
}

func (r *jsonConfigReader) Read() (*domain.Config, error) {
	return nil, errors.New("not Implemented")
}
