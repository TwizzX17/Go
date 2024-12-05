package dbSetup

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var DBConn *sql.DB

func InitDB() error {

	connStr, err := buildConnectionString()
	if err != nil {
		log.Fatalf("error building connection string: %v", err)
		return err
	}

	DBConn, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}

	if err = DBConn.Ping(); err != nil {
		return fmt.Errorf("error connecting to database: %v", err)
	}

	return nil
}

func buildConnectionString() (string, error) {
	bHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	if bHost == "" {
		log.Fatalf("DB_HOST is not set")
		return "", fmt.Errorf("DB_HOST is not set")
	}

	if dbPort == "" {
		log.Fatalf("DB_PORT is not set")
		return "", fmt.Errorf("DB_PORT is not set")
	}

	if dbUser == "" {
		log.Fatalf("DB_USER is not set")
		return "", fmt.Errorf("DB_USER is not set")
	}

	if dbPassword == "" {
		log.Fatalf("DB_PASSWORD is not set")
		return "", fmt.Errorf("DB_PASSWORD is not set")
	}

	if dbName == "" {
		log.Fatalf("DB_NAME is not set")
		return "", fmt.Errorf("DB_NAME is not set")
	}

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", bHost, dbPort, dbUser, dbPassword, dbName), nil
}

func ToString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}

func ToInt(ni sql.NullInt64) int {
	if ni.Valid {
		return int(ni.Int64)
	}
	return 0
}

func ToBool(nb sql.NullBool) bool {
	if nb.Valid {
		return nb.Bool
	}
	return false
}

func ToTime(nt sql.NullTime) *time.Time {
	if nt.Valid {
		return &nt.Time
	}
	return nil
}
