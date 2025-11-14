package config

import (
	"github.com/benyamin218118/todoService/domain"
	"github.com/benyamin218118/todoService/domain/contracts"
	"github.com/joho/godotenv"
)

type envConfigReader struct {
}

func getENVConfigReader() contracts.IConfigReader {
	return &envConfigReader{}
}

func (r *envConfigReader) Read() (*domain.Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	return &domain.Config{}, nil
}
