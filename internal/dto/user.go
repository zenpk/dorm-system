package dto

import "github.com/zenpk/dorm-system/internal/dal"

type GetUserInfoResp struct {
	CommonResp
	UserInfo *dal.User `json:"userInfo,omitempty"`
}
