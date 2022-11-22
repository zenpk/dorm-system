package dto

import "github.com/zenpk/dorm-system/internal/dal"

type UserRespGet struct {
	CommonResp
	User *dal.User `json:"user,omitempty"`
}
