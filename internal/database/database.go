package database

import (
	"database/sql"
	"fmt"
	"log"
	"tech-test/config"

	_ "github.com/lib/pq"
)

// Conn is the global database connection
var Conn *sql.DB

// InitDB initializes the database connection
func InitDB() {
	connectionInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)

	var err error
	Conn, err = sql.Open("postgres", connectionInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = Conn.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
