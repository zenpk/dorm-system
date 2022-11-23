package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zenpk/dorm-system/internal/handler"
)

// AllHandler Gin HTTP request handler
type AllHandler struct {
	building handler.Building
	dorm     handler.Dorm
	order    handler.Order
	team     handler.Team
	user     handler.User
}

var ginHandler AllHandler

// tempHandler is for testing
func tempHandler(c *gin.Context) {
	//ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	//rel := &dal.TeamUser{Id: 1}
	//err := dal.Table.TeamUser.Delete(ctx, rel)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//res, err := dal.Table.Team.FindByInnerJoinUserId(ctx, 1)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(*res)
	//fmt.Println(res.Deleted)
	//fmt.Printf("%T", *res)
}
