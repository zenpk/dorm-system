package util

import (
	"context"
	"time"
)

func ContextWithTimeout(timeout int) (context.Context, context.CancelFunc) {
	duration := time.Duration(timeout)
	return context.WithTimeout(context.Background(), duration*time.Second)
}
