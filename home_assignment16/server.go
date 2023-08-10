package main

import (
	"database/sql"
	"github.com/joho/godotenv"
	"github.com/mediocregopher/radix/v3"
	"go-course/home_assignment16/config"
	"go-course/home_assignment16/repository"
)

func initApp(db *sql.DB) (*config.Config, *repository.Storage, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, nil, err
	}

	conf := config.New()

	redis, err := radix.NewPool("tcp", "127.0.0.1:6379", 3)
	if err != nil {
		return nil, nil, err
	}

	storage := repository.NewStorage(db, redis)

	return conf, storage, nil
}
