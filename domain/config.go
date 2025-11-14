package domain

type S3Config struct {
	Endpoint   string `json:"endpoint"`
	BucketName string `json:"bucketname"`
	AccessKey  string `json:"access_key"`
	SecretKey  string `json:"secret_key"`
}

type Config struct {
	ListenHost string   `json:"listen_host"`
	ListenPort int      `json:"listen_port"`
	DBDSN      string   `json:"db_dsn"`
	RedisUrl   string   `json:"redis_url"`
	S3         S3Config `json:"s3"`
}
