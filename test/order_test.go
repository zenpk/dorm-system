package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zenpk/dorm-system/internal/controller"
	"net/http/httptest"
	"testing"
)

func BenchmarkSubmit(b *testing.B) {
	initTestEnv()
	router := gin.Default()
	controller.InitRouter(router)
	for i := 0; i < b.N; i++ {
		rec := httptest.NewRecorder()
		//orderJSON, _ := json.Marshal(orderReq)
		//req, _ := http.NewRequest(http.MethodPost, "/submit-order", bytes.NewBuffer(orderJSON))
		//req.Header.Add("Content-Type", "application/json")
		//router.ServeHTTP(rec, req)
		fmt.Println(rec.Body.String())
	}
}
