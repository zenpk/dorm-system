package util

import (
	"github.com/gin-gonic/gin"
	"github.com/zenpk/dorm-system/internal/middleware"
	"strconv"
)

// ParseId 将 string 格式的 id 转换为 int64 格式
func ParseId(idStr string) int64 {
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		middleware.Logger.Warn(err.Error())
	}
	return id
}

// QueryId 从 Context 中根据 name 提取 int64 格式的 id
func QueryId(c *gin.Context, name string) int64 {
	idStr := c.Query(name)
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		middleware.Logger.Warn(err.Error())
	}
	return id
}
