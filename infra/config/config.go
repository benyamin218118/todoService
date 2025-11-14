package config

import (
	"fmt"

	"github.com/benyamin218118/todoService/domain"
	"github.com/benyamin218118/todoService/domain/contracts"
)

type ReaderType string

const (
	ENVConfigReader  ReaderType = "env"
	JSONConfigReader ReaderType = "json"
)

func Read(rt ReaderType) (*domain.Config, error) {
	var reader contracts.IConfigReader
	switch rt {
	case ENVConfigReader:
		{
			reader = getENVConfigReader()
		}
	case JSONConfigReader:
		{
			reader = getJSONConfigReader()

		}
	default:
		{
			return nil, fmt.Errorf("invalid config reader '%s'", string(rt))
		}
	}
	conf, err := reader.Read()
	return conf, err
}
