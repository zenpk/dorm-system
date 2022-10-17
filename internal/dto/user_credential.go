package dto

import "github.com/zenpk/dorm-system/internal/dal"

// RegisterReq 用户注册请求
type RegisterReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

// RegisterLoginResp 用户注册/登录响应
type RegisterLoginResp struct {
	SuccessCode int64        `json:"success_code"`
	Status      int64        `json:"status"`
	Message     string       `json:"message"`
	Data        dal.UserInfo `json:"data"`
}

// LoginReq 用户登录请求
type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginStateResp 登录状态响应
type LoginStateResp struct {
	CurrentState int64 `json:"currentState"`
}
