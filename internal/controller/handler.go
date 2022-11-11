package controller

import "github.com/zenpk/dorm-system/internal/handler"

// HandlerSet Gin HTTP request handler
type HandlerSet struct {
	building handler.Building
	dorm     handler.Dorm
	order    handler.Order
	team     handler.Team
	user     handler.User
}

var ginHandler HandlerSet
