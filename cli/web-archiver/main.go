package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	v1 "github.com/ssst0n3/web-archiver/api/v1"
	"github.com/ssst0n3/web-archiver/archiver"
	"github.com/ssst0n3/web-archiver/config"
	"github.com/ssst0n3/web-archiver/uploader"
)

func main() {
	port := fmt.Sprintf(":%s", config.LocalListenPort)

	go archiver.Run()
	go uploader.Run()
	router := gin.Default()
	v1.Router(router)
	awesome_error.CheckFatal(router.Run(port))
}
