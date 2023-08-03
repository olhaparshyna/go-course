package main

import (
	"github.com/joho/godotenv"
	"go-course/home_assignment14/config"
	"go-course/home_assignment14/repository"
)

func initApp() (*config.Config, *repository.Processor, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, nil, err
	}

	conf := config.New()
	processor := repository.NewProcessor(&repository.Storage)

	return conf, processor, nil
}
