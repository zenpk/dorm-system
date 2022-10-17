package controller

import "github.com/zenpk/dorm-system/internal/service"

// Handler Gin HTTP request handler
type Handler struct {
	building       service.Building
	dorm           service.Dorm
	userCredential service.UserCredential
	userInfo       service.UserInfo
}

var handler Handler
