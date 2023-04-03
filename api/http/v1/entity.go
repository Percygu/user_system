package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	RequestValid    ErrCode = 0
	ErrBodyBindCode ErrCode = 1 // 参数绑定错误
	ErrParamCode    ErrCode = 2 // 请求参数不合法
)

type (
	DebugType int // debug类型
	ErrCode   int // 错误码
)

// HttpResponse http独立请求返回结构体
type HttpResponse struct {
	Code ErrCode `json:"code"`
	Msg  string  `json:"msg"`
}

// ResponseWithError http请求返回处理函数
func (rsp *HttpResponse) ResponseWithError(c *gin.Context, code ErrCode, msg string) {
	rsp.Code = code
	rsp.Msg = msg
	c.JSON(http.StatusOK, rsp)
}
