package config

import (
	"encoding/json"
	"os"

	"github.com/benyamin218118/todoService/domain"
	"github.com/benyamin218118/todoService/domain/contracts"
)

type jsonConfigReader struct {
}

func getJSONConfigReader() contracts.IConfigReader {
	return &jsonConfigReader{}
}

func (r *jsonConfigReader) Read() (conf *domain.Config, err error) {
	data, err := os.ReadFile("config.json")
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, conf); err != nil {
		return nil, err
	}
	return
}
