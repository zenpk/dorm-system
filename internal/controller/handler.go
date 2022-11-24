package controller

import (
	"github.com/zenpk/dorm-system/internal/handler"
)

// AllHandler Gin HTTP request handler
type AllHandler struct {
	dorm  handler.Dorm
	order handler.Order
	team  handler.Team
	user  handler.User
}

var ginHandler AllHandler
