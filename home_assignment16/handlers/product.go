package handlers

import (
	"encoding/json"
	"github.com/mediocregopher/radix/v3"
	"go-course/home_assignment16/repository"
	"log"
)

func ProductList(storage repository.Storage, redis radix.Pool) ([]repository.Product, error) {
	var products []repository.Product
	var jsonFromRedis []byte

	err := redis.Do(radix.Cmd(&jsonFromRedis, "GET", "products:all"))

	if err != nil {
		log.Default().Println(err)
	}

	if err == nil && len(jsonFromRedis) > 0 {
		err := json.Unmarshal(jsonFromRedis, &products)
		if err != nil {
			log.Default().Println(err)
		}

		return products, nil
	}

	products, err = storage.GetAllProducts()

	if err != nil {
		log.Default().Println(err)
		return nil, err
	}

	json, err := json.Marshal(products)
	if err != nil {
		log.Default().Println(err)
	}

	err = redis.Do(radix.FlatCmd(nil, "SET", "products:all", json))
	if err != nil {
		log.Default().Println(err)
	}

	err = redis.Do(radix.FlatCmd(nil, "EXPIRE", "products:all", 3600))
	if err != nil {
		log.Default().Println(err)
	}

	err = redis.Do(radix.Cmd(&products, "GET", "products:all"))
	if err != nil {
		log.Default().Println(err)
	}

	return products, nil
}
