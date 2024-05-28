package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lpernett/godotenv"
)

var DB *sql.DB

func Connect() (*sql.DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	// port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, host, dbname)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("error opening the db")
		return nil, err

	}
	if err := db.Ping(); err != nil {
		fmt.Println("error pinging the db")
		return nil, err
	}

	fmt.Println("Connected to db!!")
	return db, nil
}
