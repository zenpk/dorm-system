package controller

import "github.com/zenpk/dorm-system/internal/service"

// Handler Gin HTTP request handler
type Handler struct {
	userCredential service.UserCredential
	userInfo       service.UserInfo
}

var handler Handler
