package driver

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	fmt.Println("Waiting for the database to become available...")
	time.Sleep(5 * time.Second)

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging DB: %v", err)
	}

	log.Println("Connected to DB")
}

func GetDB() *sql.DB {
	return db
}

func CloseDB() {
	if db != nil {
		if err := db.Close(); err != nil {
			log.Fatalf("Error closing DB: %v", err)
		}
	}
}
