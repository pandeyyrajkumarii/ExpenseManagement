package contracts

type TransactionResponse struct {
	TxnId       string
	Amount      float64
	Category    string
	TxnType     string
	Description string
	TxnTime     string
}
