package main

import (
	router "dev11/internal/app/http"
	"dev11/internal/app/service"
	"dev11/internal/config"
	"dev11/internal/repository/postgres"
	"net/http"
)

func main() {
	cfg := config.Read()

	db := postgres.New(cfg.Postgres)
	service := service.New(db)
	server := router.NewServer(service)
	err := http.ListenAndServe(":"+cfg.Service.Port, server)
	if err != nil {
		panic(err)
	}
}
