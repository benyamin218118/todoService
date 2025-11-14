package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/benyamin218118/todoService/domain"
	"github.com/benyamin218118/todoService/domain/contracts"
	"github.com/google/uuid"
)

type s3Storage struct {
	client *s3.Client
	bucket string
	dbConn *sql.DB
}

func NewS3Storage(conf *domain.Config, dbConn *sql.DB) contracts.IStorage {

	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if conf.S3.Endpoint != "" {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           conf.S3.Endpoint,
				SigningRegion: "us-east-1",
			}, nil
		}

		// returning EndpointNotFoundError will allow the service to fallback to it's default resolution
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
		config.WithEndpointResolverWithOptions(customResolver),
	)

	if err != nil {
		panic(err)
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		if len(conf.S3.Endpoint) > 0 {
			o.UsePathStyle = true
			o.BaseEndpoint = aws.String(conf.S3.Endpoint)
		}
		o.Credentials = credentials.NewStaticCredentialsProvider(conf.S3.AccessKey, conf.S3.SecretKey, "")
	})

	s3 := s3Storage{
		client: client,
		bucket: conf.S3.BucketName,
		dbConn: dbConn,
	}
	s3.EnsureBucket(context.Background())
	return &s3
}

func (s *s3Storage) EnsureBucket(ctx context.Context) {
	_, err := s.client.CreateBucket(ctx, &s3.CreateBucketInput{
		Bucket: &s.bucket,
	})
	if err != nil {
		panic(err)
	}
}

func (s *s3Storage) Upload(file io.Reader, filename string) (*string, error) {
	id := uuid.New().String()
	input := &s3.PutObjectInput{
		Bucket: &s.bucket,
		Key:    &id,
		Body:   file,
	}

	s.dbConn.Exec(`insert into files(id,name) values(?,?)`, id, filename)
	_, err := s.client.PutObject(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to upload: %w", err)
	}
	return &id, nil
}

func (s *s3Storage) GetFileName(ctx context.Context, id string) (name *string, err error) {
	err = s.dbConn.QueryRowContext(ctx, `select name from files where id = ?`, id).Scan(&name)
	return
}
