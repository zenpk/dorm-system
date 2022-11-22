package dto

//type OrderInfo struct {
//	Id          uint64 `json:"id,omitempty"`
//	BuildingNum string `json:"buildingNum,omitempty"`
//	DormNum     string `json:"dormNum,omitempty"`
//	Info        string `json:"info,omitempty"`
//	Success     bool   `json:"success,omitempty"`
//	Deleted     bool   `json:"deleted,omitempty"`
//}

type OrderReqSubmit struct {
	BuildingNum string `json:"buildingNum,omitempty"`
}
