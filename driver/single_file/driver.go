package single_file

import (
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"io"
	"net/http"
	"net/url"
	"web-archiver/config"
)

func SingleFile(target string) (reader io.ReadCloser, err error) {
	resp, err := http.PostForm(config.SingleFileUrl, url.Values{
		"url": {target},
	})
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	reader = resp.Body
	return
}
