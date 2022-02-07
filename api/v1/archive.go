package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/lightweight_api"
	web_archiver "github.com/ssst0n3/web-archiver"
	"net/http"
)

func Archive(c *gin.Context) {
	url := c.PostForm("url")
	//if len(url) > 0 {
	//	archiver.Pool <- url
	//}
	address, err := web_archiver.Archive(url)
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
	c.String(http.StatusOK, address)
}
