package dto

import "github.com/zenpk/dorm-system/internal/dal"

type UserGetInfoResp struct {
	CommonResp
	UserInfo *dal.User `json:"userInfo,omitempty"`
}
