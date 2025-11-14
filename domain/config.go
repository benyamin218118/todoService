package domain

type Config struct {
	ListenHost string `json:"listen_host"`
	ListenPort string `json:"listen_port"`
	DBDSN      string `json:"db_dsn"`
	S3Url      string `json:"s3_url"`
	RedisUrl   string `json:"redis_url"`
}
