package dto

type CommonResp struct {
	Code int64  `json:"code" eh:"pre:CodeUncaughtError"`
	Msg  string `json:"msg" eh:"err"`
}
