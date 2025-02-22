package contracts

type TransactionResponse struct {
	TxnId       string  `json:"txn_id"`
	Amount      float64 `json:"amount"`
	Category    string  `json:"category"`
	TxnType     string  `json:"txn_type"`
	Description string  `json:"description"`
	TxnTime     string  `json:"txn_time"`
}

type MultipleTransactionResponse struct {
	Transactions []*TransactionResponse
}
