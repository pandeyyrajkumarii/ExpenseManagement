package service

import (
	"ExpenseManagement/internal/txnService/contracts"
	txnRepo "ExpenseManagement/internal/txnService/repo"
	usrRepo "ExpenseManagement/internal/userService/repo"
	"k8s.io/klog/v2"
)

type Service struct {
	TxnRepo  *txnRepo.TxnRepo
	UserRepo *usrRepo.UserRepo
}

func NewService(opt ...Opt) *Service {
	sv := &Service{}
	for _, o := range opt {
		o(sv)
	}
	return sv
}

func (s *Service) SaveTransaction(req *contracts.Transaction, userID string) (*contracts.TransactionResponse, error) {
	err := s.ValidateCreateTransactionRequest(*req)
	if err != nil {
		klog.Infof("Validation falied..", "error: ", err.Error())
		return nil, err
	}
	user, err := s.UserRepo.FindByUserId(userID)
	if err != nil && err.Error() != "user_not_found" {
		klog.Infof("DB error")
		return nil, err
	}
	if user == nil {
		klog.Infof("User not found for txn..")
		return nil, nil
	}

	txnDB, err := s.TxnRepo.Create(req, userID)
	if err != nil {
		return nil, err
	}
	klog.Infof("Txn created successfully", "Txn:- ", txnDB)
	return &contracts.TransactionResponse{
		TxnId:       txnDB.TxnId,
		Amount:      txnDB.Amount,
		Category:    txnDB.Category,
		TxnType:     txnDB.TxnType,
		Description: txnDB.Description,
		TxnTime:     txnDB.TxnTime}, nil

}
