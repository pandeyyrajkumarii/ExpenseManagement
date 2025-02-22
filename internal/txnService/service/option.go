package service

import (
	"ExpenseManagement/internal/txnService/repo"
	usrRepo "ExpenseManagement/internal/userService/repo"
	"ExpenseManagement/packages/database"
)

type Opt func(service *Service)

func WithRepo(dbInstance *database.DatabaseGorm) Opt {
	return func(service *Service) {
		service.TxnRepo = repo.NewTxnRepo(dbInstance)
	}
}

func WithUserRepo(dbInstance *database.DatabaseGorm) Opt {
	return func(service *Service) {
		service.UserRepo = usrRepo.NewUserRepo(dbInstance)
	}
}
