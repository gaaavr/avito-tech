package main

import (
	"avito/configs"
	app2 "avito/internal/app"
	"avito/internal/repository"
	"avito/internal/service"
	"fmt"
	"github.com/caarlos0/env/v6"
	"log"
)

func main() {
	var cfg configs.Common
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}
	cfg.DbPassword = "qwerty"
	fmt.Println(cfg)
	repo, _ := repository.NewPostgresDB(cfg.ConfigDB)
	r := repository.NewRepository(repo)
	s := service.NewService(r)
	app := app2.NewApp(s)
	app.Run()
}
