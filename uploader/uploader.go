package uploader

import (
	"web-archiver/config"
	"web-archiver/s3/huaweicloud"
)

type Metadata struct {
	Url      string
	FilePath string
	Content  []byte
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
	obs, err := huaweicloud.NewObs(config.AK, config.SK, config.Endpoint, config.BucketName)
	if err != nil {
		return
	}
	err = obs.Upload(metadata.Url, metadata.Content)
	if err != nil {
		return
	}
	return
}
