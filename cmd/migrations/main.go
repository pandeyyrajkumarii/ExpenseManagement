package main

import (
	"ExpenseManagement/cmd/config"
	_ "ExpenseManagement/internal/migrations"
	"ExpenseManagement/packages/database"
	"fmt"
	"github.com/pressly/goose/v3"
	"k8s.io/klog/v2"
	"log"
)

func main() {
	fmt.Println("WELCOME TO MANAGE YOUR EXPENSES....")
	cfg, err := config.NewConfig("")
	if err != nil {
		log.Fatal("error to laod config", err.Error())
	}
	klog.Infof("config is", cfg)
	err = goose.SetDialect(cfg.Database.Dialect)
	if err != nil {
		log.Fatal("error to set dialect", err.Error())
	}

	dbInstance, err := database.NewDb(&cfg.Database)
	if err != nil {
		log.Fatal("error to initiate db instance", err.Error())
	}
	sqlDb, err := dbInstance.GetInstance()
	if err != nil {
		log.Fatal("error to fetch sql db", err.Error())
	}
	err = goose.Run("up", sqlDb, "internal/migrations")
	if err != nil {
		log.Fatal("error to run migrations", err.Error())
	}

}
