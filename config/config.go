package config

import "os"

const (
	EnvLocalListenPort = "LOCAL_LISTEN_PORT"
)

var (
	LocalListenPort = os.Getenv(EnvLocalListenPort)
)

func init() {
	if len(LocalListenPort) == 0 {
		LocalListenPort = "8080"
	}
}
