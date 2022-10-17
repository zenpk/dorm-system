package dto

type GetAvailableCountResp struct {
	CommonResp
	Count int64 `json:"count"`
}

type GetAvailableBuildingsResp struct {
	CommonResp
	BuildingIds []int64 `json:"buildingIds"`
}
