package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"kirkagram/internal/config"

	_ "github.com/lib/pq"
)

type ConnectionInfo struct {
	Host     string
	Port     int
	Username string
	DBName   string
	SSLMode  string
	Password string
}

var (
	ErrUserNotFound      = errors.New("User not found")
	ErrEmailExists       = errors.New("Email already exists")
	ErrIncorrectPassword = errors.New("Incorrect password")
)

func New(cfg *config.Config) *sql.DB {
	// info := ConnectionInfo{
	// 	Host:     "localhost",
	// 	Port:     5432,
	// 	Username: "myuser",
	// 	DBName:   "mydatabase",
	// 	SSLMode:  "disable",
	// 	Password: "mypassword",
	// }

	connStr := "host=localhost port=5467 user=user password=12345 dbname=kirkagram sslmode=disable"

	// db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
	// 	info.Host, info.Port, info.Username, info.Password, info.DBName, info.SSLMode))

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		fmt.Printf("failed to connect to database: %v\n", err)
		panic(err)
	}

	return db
}
