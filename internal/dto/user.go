package dto

import "github.com/zenpk/dorm-system/internal/dal"

type GetUserInfoResp struct {
	CommonResp
	UserInfo *dal.UserInfo `json:"userInfo,omitempty"`
}
