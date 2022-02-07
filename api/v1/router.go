package v1

import (
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	router.POST("/archive", Archive)
}
