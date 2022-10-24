package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zenpk/dorm-system/internal/controller"
	pb "github.com/zenpk/dorm-system/internal/service/order"
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkSubmit(b *testing.B) {
	initTestEnv()
	router := gin.Default()
	controller.InitRouter(router)
	for i := 0; i < b.N; i++ {
		rec := httptest.NewRecorder()
		orderReq := pb.OrderRequest{
			BuildingId: 6,
			StudentId1: 90,
			StudentId2: 91,
			StudentId3: 92,
			StudentId4: 93,
		}
		orderJSON, _ := json.Marshal(orderReq)
		req, _ := http.NewRequest(http.MethodPost, "/submit-order", bytes.NewBuffer(orderJSON))
		req.Header.Add("Content-Type", "application/json")
		router.ServeHTTP(rec, req)
		fmt.Println(rec.Body.String())
	}
}
