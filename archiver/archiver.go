package archiver

import (
	"context"
	"github.com/go-shiori/obelisk"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"web-archiver/uploader"
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
	result, err := Archive(url)
	if err != nil {
		return
	}
	meta := uploader.Metadata{
		Url:     url,
		Content: result,
	}
	uploader.Pool <- meta
}

func Archive(url string) (content []byte, err error) {
	req := obelisk.Request{
		URL: url,
	}
	arc := obelisk.Archiver{EnableLog: true}
	arc.Validate()
	content, _, err = arc.Archive(context.Background(), req)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}
