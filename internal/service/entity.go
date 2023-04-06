package service

// RegisterRequest 注册请求
type RegisterRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	Gender   string `json:"gender"`
}

// LoginRequest 登陆请求
type LoginRequest struct {
	UserName string `json:"name"`
	PassWord string `json:"password"`
}

// LogoutRequest 登出请求
type LogoutRequest struct {
	UserName string
}

// GetUserInfoRequest 获取用户信息请求
type GetUserInfoRequest struct {
	UserName string
}

// GetUserInfoResponse 获取用户信息返回结构
type GetUserInfoResponse struct {
	UserName string
	Age      int
	Gender   string
	PassWord string
}

// UpdateUserInfoRequest 修改用户信息返回结构
type UpdateUserInfoRequest struct {
	UserName string // 不可更改
	Age      int
	Gender   string
	PassWord string
}
