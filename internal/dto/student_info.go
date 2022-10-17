package dto

import "github.com/zenpk/dorm-system/internal/dal"

// GetAllStudentInfosResp 获取全部学生信息响应
type GetAllStudentInfosResp struct {
	SuccessCode int64             `json:"success_code"` // 1 成功 0 失败
	Msg         string            `json:"msg"`          // 获取学生表单数据成功
	Data        []dal.StudentInfo `json:"data"`
}

// GetStudentInfoByIdResp 根据 id 获取学生信息响应
type GetStudentInfoByIdResp struct {
	SuccessCode int64           `json:"success_code"` // 1 成功 0 失败
	Msg         string          `json:"msg"`          // 添加学生数据成功
	Data        dal.StudentInfo `json:"data"`
}

// AddStudentInfoReq 添加单个学生请求
type AddStudentInfoReq struct {
	ErrorCode int64           `json:"error_code"`
	Data      dal.StudentInfo `json:"data"`
}

// AddStudentInfoResp 添加单个学生响应
type AddStudentInfoResp struct {
	SuccessCode int64           `json:"success_code"` // 1 成功 0 失败
	Msg         string          `json:"msg"`          // 添加学生数据成功
	Data        dal.StudentInfo `json:"data"`
}

// EditStudentInfoReq 修改/编辑单个学生请求
type EditStudentInfoReq struct {
	Data dal.StudentInfo `json:"data"`
}

// EditStudentInfoResp 修改/编辑单个学生响应
type EditStudentInfoResp struct {
	SuccessCode int64  `json:"success_code"` // 1 成功 0 失败
	Msg         string `json:"msg"`          // 修改学生数据成功
}

// DeleteStudentInfoReq 删除单个学生请求
type DeleteStudentInfoReq struct {
	Id int64 `json:"id"`
}

// DeleteStudentInfoResp 删除单个学生响应
type DeleteStudentInfoResp struct {
	SuccessCode int64           `json:"success_code"` // 1 成功 0 失败
	Status      int64           `json:"status"`       // 0 无数据 1 有数据
	Message     string          `json:"message"`      // 添加学生数据成功
	Data        dal.StudentInfo `json:"data"`
}
