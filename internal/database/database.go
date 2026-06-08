package database

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type DB struct {
	Conn *sql.DB
}

func NewDatabase(driver, dataSource string) (*DB, error) {
	conn, err := sql.Open(driver, dataSource)
	if err != nil {
		return nil, fmt.Errorf("urls db: %w", err)
	}

	err = conn.Ping()
	if err != nil {
		return nil, fmt.Errorf("urls db: %w", err)
	}

	return &DB{Conn: conn}, nil
}

func (db *DB) CloseDatabase() error {
	err := db.Conn.Close()
	if err != nil {
		return fmt.Errorf("urls db: %w", err)
	}
	return nil
}
