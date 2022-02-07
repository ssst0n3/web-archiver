package archiver

import (
	"web-archiver/uploader"
)

var (
	Pool = make(chan string, 10)
)

func Run() {
	for {
		url := <-Pool
		go Archive(url)
	}
}

func Archive(url string) {
	path := ""
	meta := uploader.Metadata{
		Url:      url,
		FilePath: path,
	}
	uploader.Pool <- meta
}
