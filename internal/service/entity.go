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
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
}

// LogoutRequest 登出请求
type LogoutRequest struct {
	UserName string `json:"user_name"`
}

// GetUserInfoRequest 获取用户信息请求
type GetUserInfoRequest struct {
	UserName string `json:"user_name"`
}

// GetUserInfoResponse 获取用户信息返回结构
type GetUserInfoResponse struct {
	UserName string `json:"user_name"`
	Age      int    `json:"age"`
	Gender   string `json:"gender"`
	PassWord string `json:"pass_word"`
	NickName string `json:"nick_name"`
}

// UpdateUserInfoRequest 修改用户信息返回结构
type UpdateUserInfoRequest struct {
	UserName string `json:"user_name"`
	Age      int    `json:"age"`
	Gender   string `json:"gender"`
	PassWord string `json:"pass_word"`
}
