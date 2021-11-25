package handler

import "github.com/gin-gonic/gin"

const (
	errOK           = "success"
	errParamInvalid = "param invalid"
)

func MkResponse(ctx *gin.Context, code int, msg string, data interface{}) {
	ctx.JSON(code, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}
