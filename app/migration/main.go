package main

import (
	"log"

	"github.com/diegosperes/pismo/app/model"
	"github.com/diegosperes/pismo/app/util"
)

func main() {
	deps := util.SetupApp()

	migrationErr := deps.Database.AutoMigrate(&model.Account{}, &model.Transaction{})

	if migrationErr != nil {
		log.Fatal("An error ocurred on migration; ", migrationErr.Error())
	}
}
