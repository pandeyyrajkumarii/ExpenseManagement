package repo

import (
	"ExpenseManagement/internal/txnService/contracts"
	"ExpenseManagement/internal/txnService/model"
	"ExpenseManagement/packages/database"
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

func ConvertUnixToDatetime(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format("2006-01-02 15:04:05")
}
