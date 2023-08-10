package database

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

func InitDB() *sql.DB {
	config := mysql.Config{
		User:   "user",
		Passwd: "password",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "my-app",
	}

	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err.Error())
	}

	createOrderTableQuery := `CREATE TABLE IF NOT EXISTS orders (
			id VARCHAR(255) PRIMARY KEY,
			userId INT NOT NULL,
			items TEXT, 
            FOREIGN KEY (userId) REFERENCES users (id)
		);`

	createUserTableQuery := `
		CREATE TABLE IF NOT EXISTS users (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(255),
			email VARCHAR(255) UNIQUE
		);`

	createProductTableQuery := `
		CREATE TABLE IF NOT EXISTS products (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(255),
			price DECIMAL(10, 2)
		);`

	_, err = db.Exec(createUserTableQuery)
	if err != nil {
		log.Fatal(err)
	}

	log.Default().Println("User table created or exists")

	_, err = db.Exec(createOrderTableQuery)
	if err != nil {
		log.Fatal(err)
	}

	log.Default().Println("Order table created or exists")

	_, err = db.Exec(createProductTableQuery)
	if err != nil {
		log.Fatal(err)
	}

	log.Default().Println("Product table created or exists")

	return db
}
