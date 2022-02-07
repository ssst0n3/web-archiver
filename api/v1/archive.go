package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/lightweight_api"
	"net/http"
	"web-archiver/driver/single_file"
	"web-archiver/uploader"
)

func Archive(c *gin.Context) {
	url := c.PostForm("url")
	//if len(url) > 0 {
	//	archiver.Pool <- url
	//}
	reader, err := single_file.SingleFile(url)
	if err != nil {
		return
	}
	metadata := uploader.Metadata{
		Url:     url,
		Content: reader,
	}
	address, err := uploader.Upload(metadata)
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
	c.String(http.StatusOK, address)
}
