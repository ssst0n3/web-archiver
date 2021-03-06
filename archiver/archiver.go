package archiver

import (
	"github.com/ssst0n3/web-archiver/driver/single_file"
	"github.com/ssst0n3/web-archiver/uploader"
)

var (
	Pool = make(chan string, 10)
)

func Run() {
	for {
		url := <-Pool
		go Work(url)
	}
}

func Work(url string) {
	reader, err := single_file.SingleFile(url)
	if err != nil {
		return
	}
	meta := uploader.Metadata{
		Url:     url,
		Content: reader,
	}
	uploader.Pool <- meta
}
