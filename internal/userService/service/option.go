package service

import (
	"ExpenseManagement/internal/userService/repo"
	"ExpenseManagement/packages/database"
)

type Opt func(service *Service)

func WithRepo(dbInstance *database.DatabaseGorm) Opt {
	return func(sv *Service) {
		sv.Repo = repo.NewUserRepo(dbInstance)
	}
}
