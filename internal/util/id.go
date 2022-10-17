package util

import (
	"github.com/gin-gonic/gin"
	"github.com/zenpk/dorm-system/pkg/zap"
	"strconv"
)

// ParseId string -> int64
func ParseId(idStr string) int64 {
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.Logger.Warn(err.Error())
	}
	return id
}

// QueryId extract id from GET parameters
func QueryId(c *gin.Context, name string) int64 {
	idStr := c.Query(name)
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.Logger.Warn(err.Error())
	}
	return id
}
