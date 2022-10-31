package controller

import "github.com/zenpk/dorm-system/internal/handler"

// HandlerSet Gin HTTP request handler
type HandlerSet struct {
	building       handler.Building
	userCredential handler.UserCredential
	//userInfo       handler.UserInfo
	order handler.Order
	dorm  handler.Dorm
}

var ginHandler HandlerSet
