package config

import "os"

const (
	EnvLocalListenPort = "LOCAL_LISTEN_PORT"
	EnvAK              = "AK"
	EnvSK              = "SK"
	EnvEndpoint        = "ENDPOINT"
	EnvBucketName      = "BUCKET_NAME"
)

var (
	LocalListenPort = os.Getenv(EnvLocalListenPort)
	AK              = os.Getenv(EnvAK)
	SK              = os.Getenv(EnvSK)
	Endpoint        = os.Getenv(EnvEndpoint)
	BucketName      = os.Getenv(EnvBucketName)
)

func init() {
	if len(LocalListenPort) == 0 {
		LocalListenPort = "8080"
	}
}
