package common

// CommonResponse /** Default response type that will help return customized info
// {"code": "0000", "message":"success", "error":null,"data":[]}
type CommonResponse struct {
	Code	interface{}	`json:"code"`
	Message	interface{}	`json:"message"`
	Error	interface{}	`json:"error"`
	Data	interface{}	`json:"data"`
}

// NewResponse /** Warp the response info in a object
func NewResponse(code string, data interface{}) CommonResponse {
	res := CommonResponse{}
	res.Code = code
	res.Data = data
	return res
}