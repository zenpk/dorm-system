package dto

type CommonResp struct {
	Code int32  `json:"code,omitempty" ep:"err.code"`
	Msg  string `json:"msg,omitempty" ep:"err.msg"`
}
