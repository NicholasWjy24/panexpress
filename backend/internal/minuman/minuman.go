package minuman

type Minuman struct {
	MimnID int     `json:"mimnid" gorm:"column:mimnid"`
	MimnNm string  `json:"mimnnm" gorm:"column:mimnnm"`
	MimnTp string  `json:"mimntp" gorm:"column:mimntp"`
	MimnPr float64 `json:"mimnpr" gorm:"column:mimnpr"`
	MimnQt int     `json:"mimnqt" gorm:"column:mimnqt"`
	MimnSt bool    `json:"mimnst" gorm:"column:mimnst"`
}
