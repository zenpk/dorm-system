package controller

import "github.com/zenpk/dorm-system/internal/handler"

// Handler Gin HTTP request handler
type Handler struct {
	building       handler.Building
	dorm           handler.Dorm
	userCredential handler.UserCredential
	userInfo       handler.UserInfo
}

var ginHandler Handler
