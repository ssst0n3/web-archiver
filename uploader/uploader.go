package uploader

import "github.com/ssst0n3/awesome_libs/log"

type Metadata struct {
	Url      string
	FilePath string
}

var (
	Pool = make(chan Metadata, 10)
)

func Run() {
	for {
		metadata := <-Pool
		go Upload(metadata)
	}
}

func Upload(metadata Metadata) {
	url := metadata.Url
	path := metadata.FilePath
	log.Logger.Info(url, path)
}
