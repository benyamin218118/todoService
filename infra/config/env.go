package config

import (
	"errors"
	"os"
	"strconv"

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
	if err := godotenv.Load(); err != nil {
		println(".env not found, using environment variables")
	}
	port, err := strconv.Atoi(os.Getenv("LISTEN_PORT"))
	if err != nil {
		return nil, errors.New("cant parse LISTEN_PORT env var")
	}

	conf := &domain.Config{
		ListenHost: os.Getenv("LISTEN_HOST"),
		ListenPort: port,
		DBDSN:      os.Getenv("DB_DSN"),
		RedisUrl:   os.Getenv("REDIS_URL"),
		S3: domain.S3Config{
			Endpoint:   os.Getenv("S3_URL"),
			BucketName: os.Getenv("S3_BUCKET"),
			AccessKey:  os.Getenv("S3_ACCESSKEY"),
			SecretKey:  os.Getenv("S3_SECRETKEY"),
		},
	}

	return conf, nil
}
