package config

import "os"

const (
	EnvLocalListenPort = "LOCAL_LISTEN_PORT"
	EnvAK              = "AK"
	EnvSK              = "SK"
	EnvEndpoint        = "ENDPOINT"
	EnvBucketName      = "BUCKET_NAME"
	EnvSingleFileUrl   = "SINGLE_FILE_URL"
)

var (
	LocalListenPort = os.Getenv(EnvLocalListenPort)
	AK              = os.Getenv(EnvAK)
	SK              = os.Getenv(EnvSK)
	Endpoint        = os.Getenv(EnvEndpoint)
	BucketName      = os.Getenv(EnvBucketName)
	SingleFileUrl   = os.Getenv(EnvSingleFileUrl)
)

func init() {
	if len(LocalListenPort) == 0 {
		LocalListenPort = "8080"
	}
}
