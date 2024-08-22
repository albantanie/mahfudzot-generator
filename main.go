package main

import (
	"log"

	"gorm.io/gorm/logger"

	"github.com/albantanie/mahfudzot-generator/api"
	"github.com/albantanie/mahfudzot-generator/db"
	"github.com/albantanie/mahfudzot-generator/model"
	"github.com/albantanie/mahfudzot-generator/repository"
	"github.com/albantanie/mahfudzot-generator/service"
)

func main() {
	dbInstance := db.NewDB()
	creds := model.Credential{
		Host:         "localhost",
		Username:     "faiz",
		Password:     "User123@11",
		DatabaseName: "mahfudzot",
		Port:         5432,
	}
	conn, err := dbInstance.Connect(&creds)
	if err != nil {
		panic(err)
	}

	//run database migration
	err = conn.AutoMigrate(
		// model.Mahfudzot{},
		model.Category{},
	)
	if err != nil {
		panic(err)
	}

	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	conn.Logger = logger.Default.LogMode(logger.Info) // Enable GORM logging

	mahfudzotRepo := repository.NewMahfudzotRepository(conn)

	mahfudzotService := service.NewMahfudzotService(mahfudzotRepo)

	mainAPI := api.NewAPI(*mahfudzotService)
	mainAPI.Start()

}
