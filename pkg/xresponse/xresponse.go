package xresponse

import (
	"net/http"

	"github.com/dabao-zhao/xgin/pkg/xerrors"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func newResponse(code int64, msg string, data interface{}) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, newResponse(http.StatusOK, "ok", data))
}

func Error(ctx *gin.Context, err error) {
	code, msg := xerrors.TransToCodeAndMsg(err)
	ctx.JSON(http.StatusOK, newResponse(code, msg, nil))
}
