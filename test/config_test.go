package test

import (
	"github.com/zenpk/dorm-system/internal/rpc"
	"log"
	"testing"
)

func TestViper(t *testing.T) {
	log.Println(rpc.InitClients())
}
