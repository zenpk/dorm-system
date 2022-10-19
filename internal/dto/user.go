package dto

import "github.com/zenpk/dorm-system/internal/dal"

type RegisterLoginReq struct {
	Username string `json:"username" eh:""`
	Password string `json:"password" eh:""`
}

type GetUserInfoResp struct {
	CommonResp
	UserInfo *dal.UserInfo `json:"userInfo,omitempty" eh:""`
}
