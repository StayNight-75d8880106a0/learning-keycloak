package helper

import "github.com/gin-gonic/gin"

type ResponseGlobal struct {
	ResponseCode    int         `json:"responseCode"`
	ResponseMessage string      `json:"responseMessage"`
	Data            interface{} `json:"data"`
	Error           interface{} `json:"error"`
}

func NewResponseGlobal(ctx *gin.Context, responseCode int, responseMessage string, data interface{}, err interface{}) {
	ctx.JSON(responseCode, ResponseGlobal{
		ResponseCode:    responseCode,
		ResponseMessage: responseMessage,
		Data:            data,
		Error:           err,
	})
}
