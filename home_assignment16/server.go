package main

import (
	"database/sql"
	"github.com/joho/godotenv"
	"go-course/home_assignment16/config"
	"go-course/home_assignment16/repository"
)

func initApp(db *sql.DB) (*config.Config, *repository.Storage, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, nil, err
	}

	conf := config.New()

	storage := repository.NewStorage(db)

	return conf, storage, nil
}
