package main

import (
	debtservice "debt-service/internal/debt-service"
	load "debt-service/internal/pkg/config"
	dbInit "debt-service/internal/pkg/postgres"
	repo "debt-service/internal/postgres"
	"debt-service/internal/service"
	"fmt"
	"log"
)

func main() {
	cfg, err := load.Load("./config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	db, err := dbInit.InitDB(*cfg)
	if err != nil {
		log.Fatal(err)
	}
	queries := dbInit.Queries(db)

	repos := repo.NewDebtRepo(&queries)
	service := service.NewService(repos)

	runServ := debtservice.NewRunService(service)

	fmt.Printf("Debt Service running on :%d port", cfg.DebtServicePort)
	if err := runServ.RUN(*cfg); err != nil {
		log.Fatal(err)
	}

}
