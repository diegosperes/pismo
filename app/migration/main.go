package main

import (
	"log"

	"github.com/diegosperes/pismo/app/model"
	"github.com/diegosperes/pismo/app/util"
)

func main() {
	util.SetupApp()
	migrationErr := util.GetDatabase().AutoMigrate(&model.Account{}, &model.Transaction{})

	if migrationErr != nil {
		log.Fatal("An error ocurred on migration; ", migrationErr.Error())
	}
}
