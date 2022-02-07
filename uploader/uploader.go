package uploader

import (
	"github.com/ssst0n3/web-archiver/config"
	"github.com/ssst0n3/web-archiver/s3/huaweicloud"
	"io"
	"net/url"
)

type Metadata struct {
	Url      string
	FilePath string
	Content  io.ReadCloser
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

func Canonicalize(u string) (key string) {
	return url.PathEscape(u)
}

func Upload(metadata Metadata) (archiveAddress string, err error) {
	obs, err := huaweicloud.NewObs(config.AK, config.SK, config.Endpoint, config.BucketName)
	if err != nil {
		return
	}
	key := Canonicalize(metadata.Url)
	archiveAddress, err = obs.Upload(key, metadata.Content)
	if err != nil {
		return
	}
	return
}
