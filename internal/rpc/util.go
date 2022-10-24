package rpc

import (
	"context"
	"time"
)

func createCtx(timeout int) (context.Context, context.CancelFunc) {
	duration := time.Duration(timeout)
	return context.WithTimeout(context.Background(), duration*time.Second)
}
