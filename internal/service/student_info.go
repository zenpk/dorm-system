package service

import (
	"github.com/gin-gonic/gin"
	"github.com/zenpk/dorm-system/internal/dal"
	"github.com/zenpk/dorm-system/internal/dto"
	"github.com/zenpk/dorm-system/internal/eh"
	"github.com/zenpk/dorm-system/internal/middleware"
	"github.com/zenpk/dorm-system/internal/util"
	"net/http"
)

type StudentInfo struct{}

// GetAll 获取全部学生表单信息
func (s *StudentInfo) GetAll(c *gin.Context) {
	var studentInfoDal dal.StudentInfo
	errHandler := eh.StudentInfo{C: c}
	studentInfos, err := studentInfoDal.FindAll()
	if err != nil {
		errHandler.GetAllErr(err)
		return
	}
	c.JSON(http.StatusOK, dto.GetAllStudentInfosResp{
		SuccessCode: 1,
		Msg:         "获取学生表单数据成功",
		Data:        studentInfos,
	})
}

// Add 添加一个学生信息
func (s *StudentInfo) Add(c *gin.Context) {
	var req dto.AddStudentInfoReq
	errHandler := eh.StudentInfo{C: c}
	if err := c.ShouldBindJSON(&req); err != nil {
		errHandler.AddErr(err)
		return
	}
	studentInfo := req.Data
	if err := studentInfo.Create(&studentInfo); err != nil {
		errHandler.AddErr(err)
		return
	}
	c.JSON(http.StatusOK, dto.AddStudentInfoResp{
		SuccessCode: 1,
		Msg:         "添加学生数据成功",
		Data:        studentInfo,
	})
}

// Edit 修改/编辑单个学生信息
func (s *StudentInfo) Edit(c *gin.Context) {
	var req dto.EditStudentInfoReq
	errHandler := eh.StudentInfo{C: c}
	if err := c.ShouldBindJSON(&req); err != nil {
		errHandler.EditErr(err)
		return
	}
	studentInfo := req.Data
	if err := studentInfo.UpdateNewFields(&studentInfo); err != nil {
		errHandler.EditErr(err)
		return
	}
	c.JSON(http.StatusOK, dto.EditStudentInfoResp{
		SuccessCode: 1,
		Msg:         "修改学生数据成功",
	})
}

// Delete 删除学生信息
func (s *StudentInfo) Delete(c *gin.Context) {
	var studentInfo dal.StudentInfo
	var req dto.DeleteStudentInfoReq
	errHandler := eh.StudentInfo{C: c}
	if err := c.ShouldBindJSON(&req); err != nil {
		errHandler.DeleteErr(err)
		return
	}
	studentInfo.StudentId = req.Id
	if err := studentInfo.Delete(&studentInfo); err != nil {
		middleware.Logger.Warn(err.Error())
		c.JSON(http.StatusOK, dto.DeleteStudentInfoResp{
			SuccessCode: 0,
			Status:      0,
			Message:     err.Error(),
			Data:        studentInfo,
		})
		return
	}
	c.JSON(http.StatusOK, dto.DeleteStudentInfoResp{
		SuccessCode: 204,
		Status:      1,
		Message:     "该学生数据已删除",
		Data:        studentInfo,
	})
}

// GetAllByParams 根据性别、机构、状态筛选用户
func (s *StudentInfo) GetAllByParams(c *gin.Context) {
	gender := c.Query("gender")
	organization := c.Query("organization")
	state := c.Query("state")
	var studentInfos []dal.StudentInfo
	var studentInfoDal dal.StudentInfo
	errHandler := eh.StudentInfo{C: c}
	studentInfos, err := studentInfoDal.FindAllByParams(gender, organization, state)
	if err != nil {
		errHandler.GetAllByParamsErr(err)
		return
	}
	c.JSON(http.StatusOK, dto.GetAllStudentInfosResp{
		SuccessCode: 1,
		Msg:         "获取学生表单数据成功",
		Data:        studentInfos,
	})
}

// GetById 根据 id 获取学生信息
func (s *StudentInfo) GetById(c *gin.Context) {
	id := util.QueryId(c, "id")
	var studentInfo dal.StudentInfo
	errHandler := eh.StudentInfo{C: c}
	studentInfo, err := studentInfo.FindByStudentId(id)
	if err != nil {
		errHandler.GetByIdErr(err)
		return
	}
	c.JSON(http.StatusOK, dto.GetStudentInfoByIdResp{
		SuccessCode: 1,
		Msg:         "获取学生表单数据成功",
		Data:        studentInfo,
	})
}
