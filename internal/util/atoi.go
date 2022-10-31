package util

import (
	"github.com/gin-gonic/gin"
	"github.com/zenpk/dorm-system/pkg/zap"
	"strconv"
)

// ParseU64 string -> uint64
func ParseU64(idStr string) uint64 {
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		zap.Logger.Warn(err.Error())
		return 0
	}
	return id
}

func ParseU32(idStr string) uint32 {
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		zap.Logger.Warn(err.Error())
		return 0
	}
	return uint32(id)
}

func Parse64(idStr string) int64 {
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.Logger.Warn(err.Error())
		return 0
	}
	return id
}

func Parse32(idStr string) int32 {
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		zap.Logger.Warn(err.Error())
		return 0
	}
	return int32(id)
}

// QueryU64 extract id from GET parameters
func QueryU64(c *gin.Context, name string) uint64 {
	idStr := c.Query(name)
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		zap.Logger.Warn(err.Error())
		return 0
	}
	return id
}

func QueryU32(c *gin.Context, name string) uint32 {
	idStr := c.Query(name)
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		zap.Logger.Warn(err.Error())
		return 0
	}
	return uint32(id)
}

func Query64(c *gin.Context, name string) int64 {
	idStr := c.Query(name)
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.Logger.Warn(err.Error())
		return 0
	}
	return id
}

func Query32(c *gin.Context, name string) int32 {
	idStr := c.Query(name)
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		zap.Logger.Warn(err.Error())
		return 0
	}
	return int32(id)
}
