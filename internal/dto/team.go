package dto

type TeamUserInfo struct {
	StudentNum string `json:"studentNum,omitempty"`
	Name       string `json:"name,omitempty"`
}

type TeamRespCreateGet struct {
	CommonResp
	Id      uint64          `json:"id,omitempty"`
	Code    string          `json:"code,omitempty"`
	Gender  string          `json:"gender,omitempty"`
	Owner   *TeamUserInfo   `json:"owner,omitempty"`
	Members []*TeamUserInfo `json:"members,omitempty"`
}
