package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	CodeSuccess           ErrCode = 0     // http请求成功
	CodeBodyBindErr       ErrCode = 10001 // 参数绑定错误
	CodeParamErr          ErrCode = 10002 // 请求参数不合法
	CodeRegisterErr       ErrCode = 10003 // 注册错误
	CodeLoginErr          ErrCode = 10003 // 登录错误
	CodeLogoutErr         ErrCode = 10004 // 登出错误
	CodeGetUserInfoErr    ErrCode = 10005 // 获取用户信息错误
	CodeUpdateUserInfoErr ErrCode = 10006 // 更新用户信息错误
)

type (
	DebugType int // debug类型
	ErrCode   int // 错误码
)

// HttpResponse http独立请求返回结构体
type HttpResponse struct {
	Code ErrCode     `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// ResponseWithError http请求返回处理函数
func (rsp *HttpResponse) ResponseWithError(c *gin.Context, code ErrCode, msg string) {
	rsp.Code = code
	rsp.Msg = msg
	c.JSON(http.StatusInternalServerError, rsp)
}

func (rsp *HttpResponse) ResponseSuccess(c *gin.Context) {
	rsp.Code = CodeSuccess
	rsp.Msg = "success"
	c.JSON(http.StatusOK, rsp)
}

func (rsp *HttpResponse) ResponseWithData(c *gin.Context, data interface{}) {
	rsp.Code = CodeSuccess
	rsp.Msg = "success"
	rsp.Data = data
	c.JSON(http.StatusOK, rsp)
}
