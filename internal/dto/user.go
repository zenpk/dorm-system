package dto

import "github.com/zenpk/dorm-system/internal/dal"

type RegisterLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterLoginResp struct {
	Code     int64        `json:"code"`
	Msg      string       `json:"msg"`
	UserInfo dal.UserInfo `json:"userInfo; omitempty"`
}
