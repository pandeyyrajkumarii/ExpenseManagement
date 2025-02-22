package repo

import (
	"ExpenseManagement/internal/txnService/contracts"
	"ExpenseManagement/internal/txnService/model"
	"ExpenseManagement/packages/database"
	"fmt"
	"k8s.io/klog/v2"
	"time"
)

type TxnRepo struct {
	Dbs *database.DatabaseGorm
}

func NewTxnRepo(db *database.DatabaseGorm) *TxnRepo {
	return &TxnRepo{Dbs: db}
}

func (t *TxnRepo) Create(txn *contracts.Transaction, userID string) (*model.TxnDb, error) {
	currentTime := time.Now().Unix()
	txnDb := ToTxnDBModel(txn, userID)
	txnDb.CreatedAt = currentTime
	txnDb.UpdatedAt = currentTime
	_, err := t.Dbs.Create(&txnDb)
	if err != nil {
		return nil, err
	}
	return txnDb, nil
}

func (t *TxnRepo) GetByQuery(req *contracts.GetTransactionRequest, userID string) (*contracts.MultipleTransactionResponse, error) {
	var transactions []*model.TxnDb
	query := t.Dbs.Db.Where("txn_time >= ? AND txn_time <= ? AND user_id = ?",
		ConvertUnixToDatetime(req.TxnFrom), ConvertUnixToDatetime(req.TxnTo), userID).Order("txn_time ASC")
	rows, err := t.Dbs.FindWithQueryFilter(&transactions, query)
	if err != nil {
		return nil, err
	}
	if rows == 0 {
		return nil, fmt.Errorf("no transactions found")
	}
	var apiTransaction []*contracts.TransactionResponse
	for _, txn := range transactions {
		apiTransaction = append(apiTransaction, ToApiResponse(txn))
	}
	klog.Infof("apiTransaction: %v", transactions[0])
	return &contracts.MultipleTransactionResponse{
		Transactions: apiTransaction,
	}, nil
}

func ToTxnDBModel(txn *contracts.Transaction, userID string) *model.TxnDb {

	return &model.TxnDb{
		TxnId:       txn.TxnId,
		UserId:      userID,
		Amount:      txn.Amount,
		Category:    txn.Category,
		TxnType:     txn.TxnType,
		TxnTime:     ConvertUnixToDatetime(txn.TxnTime),
		Description: txn.Description,
	}
}

func ToApiResponse(txn *model.TxnDb) *contracts.TransactionResponse {

	return &contracts.TransactionResponse{
		TxnId:       txn.TxnId,
		Amount:      txn.Amount,
		Category:    txn.Category,
		TxnType:     txn.TxnType,
		Description: txn.Description,
		TxnTime:     txn.TxnTime,
	}
}

func ConvertUnixToDatetime(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format("2006-01-02 15:04:05")
}
