package model

type TxnDb struct {
	TxnId       string  `gorm:"column:txn_id;not null" json:"txn_id"`
	UserId      string  `gorm:"column:user_id;not null" json:"user_id"`
	Amount      float64 `gorm:"not null" json:"amount"`
	Category    string  `gorm:"default null" json:"category"`
	TxnType     string  `gorm:"column:txn_type;not null" json:"txn_type"`
	Description string  `gorm:"default null" json:"description"`
	TxnTime     string  `gorm:"column:txn_time;not null" json:"txn_time"`
	CreatedAt   int64   `gorm:"not null" json:"created_at"`
	UpdatedAt   int64   `gorm:"not null" json:"updated_at"`
}

func (TxnDb) TableName() string {
	return "transactions"
}
