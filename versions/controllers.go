package versions

import (
	"github.com/gin-gonic/gin"
	"golearn/m/v1/common"
	"net/http"
)

func VersionRetrieve(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		VersionRetrieves(c)
		return
	}
	versionModel,err := FindOneVersionById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, common.NewResponse("0100", nil))
		return
	}

	versionSerializer := VersionSerializer{c, versionModel}
	c.JSON(http.StatusOK, common.NewResponse("0000",versionSerializer.Response()))
}

func VersionRetrieves(c *gin.Context) {
	versionModels, err := AllVersion()

	if err != nil {
		c.JSON(http.StatusNotFound, common.NewResponse("0100", nil))
		return
	}

	var data []VersionResponse
	for _,version := range versionModels  {
		VersionSerializer := VersionSerializer{c, version}
		data = append(data,VersionSerializer.Response())
	}
	c.JSON(http.StatusOK, common.NewResponse("0000",data))
}
