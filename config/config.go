package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type App struct {
	DB     *sql.DB
	JWTKey []byte
}

func InitDB() *sql.DB {
	connStr := os.Getenv("XATA_PSQL_URL")
	if connStr == "" {
		log.Fatal("XATA_PSQL_URL environment variable is not set")
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func GetJWTKey() []byte {
	key := []byte(os.Getenv("JWT_SECRET"))
	if len(key) == 0 {
		log.Fatal("JWT_SECRET environment variable is not set")
	}
	return key
}

func LoadSchema(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
