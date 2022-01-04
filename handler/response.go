package handler

import "github.com/sanscope/apk_analysis_cloud_platform_server/enum"

type Response struct {
	Code    uint16      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// the factory pattern.
func ResponseFactory(code enum.OperationStatus, data interface{}) Response {
	res := Response{
		Code:    uint16(code),
		Message: code.String(),
		Data:    data,
	}
	return res
}
