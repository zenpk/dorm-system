package controller

import "github.com/zenpk/dorm-system/internal/service"

// Handler 请求处理器
type Handler struct {
	studentInfo    service.StudentInfo
	userCredential service.UserCredential
	userInfo       service.UserInfo
}

var handler Handler
