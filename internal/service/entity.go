package service

// RegisterRequest 注册请求
type RegisterRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	Gender   string `json:"gender"`
}

// UserRequest 登录请求
type UserRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
