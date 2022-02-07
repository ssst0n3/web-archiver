package obelisk

import (
	"context"
	"github.com/go-shiori/obelisk"
	"github.com/ssst0n3/awesome_libs/awesome_error"
)

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
