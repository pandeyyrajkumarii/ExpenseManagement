package contracts

type Transaction struct {
	TxnId       string
	UserId      string
	Amount      float64
	Category    string
	TxnType     string
	Description string
	TxnTime     int64
}
