package web_archiver

import (
	"github.com/ssst0n3/web-archiver/driver/single_file"
	"github.com/ssst0n3/web-archiver/uploader"
)

func Archive(url string) (address string, err error) {
	reader, err := single_file.SingleFile(url)
	if err != nil {
		return
	}
	metadata := uploader.Metadata{
		Url:     url,
		Content: reader,
	}
	address, err = uploader.Upload(metadata)
	if err != nil {
		return
	}
	return
}
