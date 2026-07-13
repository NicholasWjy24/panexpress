package makanan

type Makanan struct {
	MknnID int     `json:"mknnid" gorm:"column:mknnid"`
	MknnNm string  `json:"mknnnm" gorm:"column:mknnnm"`
	MknnTp string  `json:"mknntp" gorm:"column:mknntp"`
	MknnPr float64 `json:"mknnpr" gorm:"column:mknnpr"`
	MknnQt int     `json:"mknnqt" gorm:"column:mknnqt"`
	MknnSt bool    `json:"mknnst" gorm:"column:mknnst"`
	MknnIg []byte  `json:"-" gorm:"column:mknnig;type:bytea"`
}
