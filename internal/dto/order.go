package dto

type OrderRequest struct {
	BuildingNum string `json:"buildingNum,omitempty"`
	StudentNum1 string `json:"studentNum1,omitempty"`
	StudentNum2 string `json:"studentNum2,omitempty"`
	StudentNum3 string `json:"studentNum3,omitempty"`
	StudentNum4 string `json:"studentNum4,omitempty"`
}
