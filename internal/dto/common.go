package dto

type CommonResp struct {
	Code int64  `json:"code" eh:"CodeUncaughtError"`
	Msg  string `json:"msg" eh:"err"`
}
