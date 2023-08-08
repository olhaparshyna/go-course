package repository

import (
	"database/sql"
	"fmt"
	"log"
)

type User struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (storage *Storage) FindUserIdByEmail(email string) int64 {
	query := "SELECT id FROM users WHERE email = ?"
	row := storage.DB.QueryRow(query, email)

	var id int64
	err := row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("User not found")
		} else {
			log.Fatal(err)
		}
		return 0
	}

	return id
}

func (storage *Storage) UserStore(name string, email string) (error, int64) {
	insertQuery := "INSERT INTO users (name, email) VALUES (?, ?)"

	result, err := storage.DB.Exec(insertQuery, name, email)
	if err != nil {
		log.Fatal(err)
		return err, 0
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return err, 0
	}
	return err, id
}
