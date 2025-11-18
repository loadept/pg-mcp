package config

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

var (
	onceDB   sync.Once
	instance *postgres
)

type postgres struct {
	db *sql.DB
}

func (s *postgres) Connect() error {
	pgURI := GetEnv("POSTGRES_URI")
	if pgURI == "" {
		return fmt.Errorf("POSTGRES_URI environment variable is required")
	}

	db, err := sql.Open("postgres", pgURI)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	s.db = db
	return nil
}

func (s *postgres) getNow() (string, error) {
	var now string

	if err := s.db.QueryRow("SELECT CURRENT_DATE::VARCHAR").Scan(&now); err != nil {
		return "", err
	}

	return now, nil
}

func NewDBPostgres() (*postgres, error) {
	var err error

	onceDB.Do(func() {
		instance = &postgres{}
		if err = instance.Connect(); err == nil {
			var now string

			now, err = instance.getNow()
			if err == nil {
				log.Printf("Connected to PostgreSQL database, current date %s\n", now)
			}
		}
	})
	return instance, err
}

func (s *postgres) GetDB() *sql.DB {
	return s.db
}

func (s *postgres) Close() error {
	return s.db.Close()
}
