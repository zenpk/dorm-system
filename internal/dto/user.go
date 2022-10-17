package dto

import "github.com/zenpk/dorm-system/internal/dal"

type RegisterLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type GetUserInfoResp struct {
	CommonResp
	UserInfo dal.UserInfo `json:"userInfo,omitempty"`
}
