package versions

import (
	"github.com/gin-gonic/gin"
)

func VersionRegister(router *gin.RouterGroup) {
	router.GET("/version", VersionRetrieve)
}