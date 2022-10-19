package dto

type GetAvailableCountResp struct {
	CommonResp
	Count int64 `json:"count" eh:"0"`
}

type GetAvailableBuildingsResp struct {
	CommonResp
	BuildingIds []int64 `json:"buildingIds" eh:""`
}
