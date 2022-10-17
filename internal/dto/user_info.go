package dto

import "github.com/zenpk/dorm-system/internal/dal"

// UserInfoFindAllResp UserInfo FindAll 返回的响应
type UserInfoFindAllResp struct {
	UserInfos []dal.UserInfo `json:"UserInfos"`
}

// UserInfoResp 用户信息响应
type UserInfoResp struct {
	UserInfo dal.UserInfo `json:"UserInfo"`
}
