package contracts

type Transaction struct {
	TxnId       string  `json:"txn_id"`
	Amount      float64 `json:"amount"`
	Category    string  `json:"category"`
	TxnType     string  `json:"txn_type"`
	Description string  `json:"description"`
	TxnTime     int64   `json:"txn_time"`
}
