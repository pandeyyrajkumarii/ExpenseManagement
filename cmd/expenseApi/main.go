package main

import (
	"ExpenseManagement/cmd/config"
	txSvc "ExpenseManagement/internal/txnService"
	svc1 "ExpenseManagement/internal/txnService/service"
	usrSvc "ExpenseManagement/internal/userService"
	svc "ExpenseManagement/internal/userService/service"
	"ExpenseManagement/packages/database"
	"ExpenseManagement/packages/server"
	"log"
)

func main() {
	cfg, err := config.NewConfig("")
	if err != nil {
		log.Fatal("error to laod config", err.Error())
	}

	dbInstance, err := database.NewDb(&cfg.Database)
	if err != nil {
		log.Fatal("error to initiate db instance", err.Error())
	}

	userSvc := svc.NewService(
		svc.WithRepo(dbInstance))
	usrSvr := usrSvc.NewServer(userSvc)
	txnSvc := svc1.NewService(
		svc1.WithRepo(dbInstance),
		svc1.WithUserRepo(dbInstance))
	txnServer := txSvc.NewServer(txnSvc)
	svr := server.NewServer(
		server.WithPort("8000"),
		server.WithUserServer(usrSvr),
		server.WithTxnServer(txnServer))
	svr.Start()
}
