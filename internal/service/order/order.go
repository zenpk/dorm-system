package order

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/zenpk/dorm-system/internal/dal"
	"github.com/zenpk/dorm-system/pkg/ep"
)

func Submit(message *sarama.ConsumerMessage) error {
	var req OrderRequest
	if err := json.Unmarshal(message.Value, &req); err != nil {
		return err
	}
	order := &dal.Order{}
	//student1, err := dal.Table.User.FindByStudentNum(nil, req.StudentNum2)
	//if err != nil {
	//	return err
	//}
	//if err := checkStudent(student1.Gender, student1); err != nil {
	//	return err
	//}
	//order.StudentId1 = student1.Id
	//if len(req.StudentNum2) > 0 {
	//	student2, err := dal.Table.User.FindByStudentNum(nil, req.StudentNum2)
	//	if err != nil {
	//		return err
	//	}
	//	if err := checkStudent(student1.Gender, student2); err != nil {
	//		return err
	//	}
	//	order.StudentId2 = student2.Id
	//}
	//if len(req.StudentNum3) > 0 {
	//	student3, err := dal.Table.User.FindByStudentNum(nil, req.StudentNum3)
	//	if err != nil {
	//		return err
	//	}
	//	if err := checkStudent(student1.Gender, student3); err != nil {
	//		return err
	//	}
	//	order.StudentId3 = student3.Id
	//}
	//if len(req.StudentNum4) > 0 {
	//	student4, err := dal.Table.User.FindByStudentNum(nil, req.StudentNum4)
	//	if err != nil {
	//		return err
	//	}
	//	if err := checkStudent(student1.Gender, student4); err != nil {
	//		return err
	//	}
	//	order.StudentId4 = student4.Id
	//}
	return dal.Table.Order.Create(nil, order)
	// TODO
}

func checkStudent(teamGender string, student *dal.User) error {
	if student.DormId != 0 {
		errPack := ep.ErrInputBody
		errPack.Msg = "someone already has a dorm"
		return errPack
	}
	if student.Gender != teamGender {
		errPack := ep.ErrInputBody
		errPack.Msg = "genders must be the same"
		return errPack
	}
	return nil
}

//func (s *Server) Submit(ctx context.Context, req *OrderRequest) (*OrderReply, error) {
//	ids := make([]uint64, 0)
//	ids = append(ids, req.StudentId1)
//	if req.StudentId2 != 0 {
//		ids = append(ids, req.StudentId2)
//	}
//	if req.StudentId3 != 0 {
//		ids = append(ids, req.StudentId3)
//	}
//	if req.StudentId4 != 0 {
//		ids = append(ids, req.StudentId4)
//	}
//	// check if already has dorm
//	users := make([]*dal.User, 0)
//	userTable := new(dal.User)
//	var teamGender string
//	for i, id := range ids {
//		user, err := userTable.FindById(ctx, id)
//		if err != nil {
//			return nil, err
//		}
//		if user.DormId != 0 {
//			errPack := ep.ErrInputBody
//			errPack.Msg = "someone already has a dorm"
//			return nil, errPack
//		}
//		users = append(users, user)
//		// check gender
//		if i == 0 {
//			teamGender = user.Gender
//		} else {
//			if user.Gender != teamGender {
//				errPack := ep.ErrInputBody
//				errPack.Msg = "genders must be the same"
//				return nil, errPack
//			}
//		}
//	}
//	num := len(users)
//	dormTable := new(dal.Dorm)
//	dorm, err := dormTable.Allocate(ctx, req.BuildingId, uint64(num), teamGender)
//	if err != nil {
//		return nil, err
//	}
//	for i := range users {
//		users[i].DormId = dorm.Id
//		if err := userTable.Update(ctx, users[i]); err != nil {
//			return nil, err
//		}
//	}
//	order := &dal.Order{
//		DormId:     dorm.Id,
//		StudentId1: req.StudentId1,
//		StudentId2: req.StudentId2,
//		StudentId3: req.StudentId3,
//		StudentId4: req.StudentId4,
//	}
//	if err := order.Create(ctx, order); err != nil {
//		return nil, err
//	}
//	reply := &OrderReply{
//		Resp: &common.CommonResponse{
//			Code: ep.ErrOK.Code,
//			Msg:  "submit success",
//		},
//	}
//	return reply, nil
//}
