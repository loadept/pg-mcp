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

// postgres encapsulates a PostgreSQL database connection.
// It provides methods for connection management and query execution.
type postgres struct {
	db *sql.DB
}

// Connect establishes a connection to the PostgreSQL database.
// It receives the connection URI as a parameter, configures
// connection pool settings, and validates the connection with a ping.
//
// Parameters:
//   - pgURI: PostgreSQL connection URI string
//
// Returns:
//   - error: Any error encountered during connection establishment
func (s *postgres) Connect(pgURI string) error {
	if pgURI == "" {
		return fmt.Errorf("postgres URI is required, please provide it as a parameter")
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

// getNow retrieves the current date from the database.
// It executes a simple query to verify database connectivity and
// returns the current date as a string.
//
// Returns:
//   - string: Current database date in string format
//   - error: Any error encountered during query execution
func (s *postgres) getNow() (string, error) {
	var now string

	if err := s.db.QueryRow("SELECT CURRENT_DATE::VARCHAR").Scan(&now); err != nil {
		return "", err
	}

	return now, nil
}

// NewDBPostgres creates a new PostgreSQL connection instance.
// It implements the singleton pattern to ensure only one database connection
// is created. The function is safe for concurrent use.
//
// Returns:
//   - *postgres: Singleton postgres instance with established connection
//   - error: Any error encountered during connection initialization
func NewDBPostgres(pgURI string) (*postgres, error) {
	var err error

	onceDB.Do(func() {
		instance = &postgres{}
		if err = instance.Connect(pgURI); err == nil {
			var now string

			now, err = instance.getNow()
			if err == nil {
				log.Printf("Connected to PostgreSQL database, current date %s\n", now)
			}
		}
	})
	return instance, err
}

// GetDB returns the underlying database connection.
// It provides access to the raw *sql.DB for query execution.
//
// Returns:
//   - *sql.DB: The database connection pool
func (s *postgres) GetDB() *sql.DB {
	return s.db
}

// Close closes the database connection.
// It should be called when the application is shutting down to properly
// release database resources.
//
// Returns:
//   - error: Any error encountered during connection closure
func (s *postgres) Close() error {
	return s.db.Close()
}
