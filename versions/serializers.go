package versions

import "github.com/gin-gonic/gin"

type VersionSerializer struct {
	C *gin.Context
	VersionModel
}

type VersionResponse struct {
	App			string  `json:"app"`
	Version		string  `json:"version"`
	Code		*int	`json:"code"`
}

func (s *VersionSerializer) Response() VersionResponse {
	version := VersionResponse{
		App:		s.App,
		Version:	s.Version,
		Code:       s.Code,
	}
	return version
}

