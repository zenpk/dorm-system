package eh

import (
	"github.com/gin-gonic/gin"
	"github.com/zenpk/dorm-system/internal/dal"
	"github.com/zenpk/dorm-system/internal/dto"
	"github.com/zenpk/dorm-system/internal/middleware"
	"net/http"
)

type StudentInfo struct {
	C *gin.Context
}

// GetAllErr 获取学生错误，自动记录错误日志并返回给前端
func (s *StudentInfo) GetAllErr(err error) {
	middleware.Logger.Warn(err.Error())
	s.C.JSON(http.StatusOK, dto.GetAllStudentInfosResp{
		SuccessCode: 0,
		Msg:         err.Error(),
		Data:        nil,
	})
}

// AddErr 添加学生错误，自动记录错误日志并返回给前端
func (s *StudentInfo) AddErr(err error) {
	middleware.Logger.Warn(err.Error())
	s.C.JSON(http.StatusOK, dto.AddStudentInfoResp{
		SuccessCode: 0,
		Msg:         err.Error(),
		Data:        dal.StudentInfo{},
	})
}

func (s *StudentInfo) EditErr(err error) {
	middleware.Logger.Warn(err.Error())
	s.C.JSON(http.StatusOK, dto.EditStudentInfoResp{
		SuccessCode: 0,
		Msg:         err.Error(),
	})
}

func (s *StudentInfo) DeleteErr(err error) {
	middleware.Logger.Warn(err.Error())
	middleware.Logger.Warn(err.Error())
	s.C.JSON(http.StatusOK, dto.DeleteStudentInfoResp{
		SuccessCode: 0,
		Status:      0,
		Message:     err.Error(),
		Data:        dal.StudentInfo{},
	})
}

func (s *StudentInfo) GetAllByParamsErr(err error) {
	middleware.Logger.Warn(err.Error())
	s.C.JSON(http.StatusOK, dto.GetAllStudentInfosResp{
		SuccessCode: 0,
		Msg:         err.Error(),
		Data:        nil,
	})
}

func (s *StudentInfo) GetByIdErr(err error) {
	middleware.Logger.Warn(err.Error())
	s.C.JSON(http.StatusOK, dto.GetStudentInfoByIdResp{
		SuccessCode: 0,
		Msg:         err.Error(),
		Data:        dal.StudentInfo{},
	})
}
