package dto

type CommonResp struct {
	Code int32  `json:"code" ep:"err.code"`
	Msg  string `json:"msg" ep:"err.msg"`
}
