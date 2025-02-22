package model

type TxnDb struct {
	TxnId       string  `gorm:"column:txnid;not null" json:"txnId"`
	UserId      string  `gorm:"column:userid;not null" json:"userid"`
	Amount      float64 `gorm:"not null" json:"amount"`
	Category    string  `gorm:"default null" json:"category"`
	TxnType     string  `gorm:"column:transactionType;not null" json:"txnType"`
	Description string  `gorm:"default null" json:"description"`
	TxnTime     string  `gorm:"column:txntime;not null" json:"txnTime"`
	CreatedAt   int64   `gorm:"not null" json:"createdAt"`
	UpdatedAt   int64   `gorm:"not null" json:"updatedAt"`
}

func (TxnDb) TableName() string {
	return "tnxs"
}
