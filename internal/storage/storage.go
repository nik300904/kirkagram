package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	internalConfig "kirkagram/internal/config"

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
	ErrUserNotFound              = errors.New("User not found")
	ErrEmailAlreadyRegistered    = errors.New("User with this email already exists")
	ErrUsernameAlreadyRegistered = errors.New("User with this username already exists")
	ErrUserAlreadyExists         = errors.New("User already exists")
	ErrIncorrectPassword         = errors.New("Incorrect password")
	ErrNoSuchKey                 = errors.New("No such key")
	ErrPostExists                = errors.New("Post already exists")
	ErrPostNotFound              = errors.New("Post not found")
	ErrPostAlreadyLiked          = errors.New("Post already liked")
)

func New(cfg *internalConfig.Config) *sql.DB {
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

func NewS3Client() *s3.Client {
	cfgS3, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		panic(err)
	}

	client := s3.NewFromConfig(cfgS3)

	return client
}
