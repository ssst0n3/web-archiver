package v1

import (
	"github.com/gin-gonic/gin"
	"web-archiver/archiver"
)

func Archive(c *gin.Context) {
	url := c.PostForm("url")
	if len(url) > 0 {
		archiver.Pool <- url
	}
}
