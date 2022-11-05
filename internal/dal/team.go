package dal

type Team struct {
	Id          uint64 `gorm:"primaryKey" json:"-"`
	StudentNum1 string `gorm:"not null; unique" json:"studentId1,omitempty"`
	StudentNum2 string `json:"studentId2,omitempty"`
	StudentNum3 string `json:"studentId3,omitempty"`
	StudentNum4 string `json:"studentId4,omitempty"`
}
