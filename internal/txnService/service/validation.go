package service

import (
	"ExpenseManagement/internal/txnService/contracts"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (s *Service) ValidateCreateTransactionRequest(t contracts.Transaction) error {
	return validation.ValidateStruct(&t,
		validation.Field(&t.TxnId, validation.Required, validation.Length(5, 255)),
		validation.Field(&t.Amount, validation.Required, validation.Min(0.01)),
		validation.Field(&t.TxnType, validation.Required, validation.In("debit", "credit")),
		validation.Field(&t.TxnTime, validation.Required),
	)
}
