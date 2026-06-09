package database

import (
	"database/sql"
	"fmt"
	"time"
)

type responseURL struct {
	ID        int64  `json:"id"`
	URL       string `json:"url"`
	ShortCode string `json:"shortCode"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func (db *DB) CreateShortURL(longURL, shortCode string) (int64, error) {
	stmt, err := db.Conn.Prepare("INSERT INTO urls (url, short_code) VALUES (?, ?)")
	if err != nil {
		return 0, fmt.Errorf("CreateShortURL: %w", err)
	}
	result, err := stmt.Exec(longURL, shortCode)
	if err != nil {
		return 0, fmt.Errorf("CreateShortURL: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("CreateShortURL: %w", err)
	}
	return id, nil
}

func (db *DB) UpdateOriginalByShortURL(longURL, shortCode string) error {
	stmt, err := db.Conn.Prepare("UPDATE urls SET url = ?, updated_at = ? WHERE short_code = ?")
	if err != nil {
		return fmt.Errorf("OriginalByShortURL: %w", err)
	}
	result, err := stmt.Exec(longURL, time.Now().UTC(), shortCode)
	if err != nil {
		return fmt.Errorf("OriginalByShortURL: %w", err)
	}
	num, err := result.RowsAffected()
	if num == 0 {
		return sql.ErrNoRows
	}
	if err != nil {
		return fmt.Errorf("OriginalByShortURL: %w", err)
	}
	return nil
}

func (db *DB) DeleteShortURL(shortCode string) error {
	stmt, err := db.Conn.Prepare("DELETE FROM urls WHERE short_code = ?")
	if err != nil {
		return fmt.Errorf("DeleteShortURL: %w", err)
	}
	result, err := stmt.Exec(shortCode)
	if err != nil {
		return fmt.Errorf("DeleteShortURL: %w", err)
	}
	num, err := result.RowsAffected()
	if num == 0 {
		return sql.ErrNoRows
	}
	if err != nil {
		return fmt.Errorf("DeleteShortURL: %w", err)
	}
	return nil
}

func (db *DB) GetShortURLByID(id int64) (*responseURL, error) {
	var result responseURL

	row := db.Conn.QueryRow("SELECT id, url, short_code, created_at, updated_at FROM urls WHERE id = ?", id)
	if err := row.Scan(&result.ID, &result.URL, &result.ShortCode, &result.CreatedAt, &result.UpdatedAt); err != nil {
		return nil, fmt.Errorf("ShortURLsByID %d: %w", id, err)
	}

	return &result, nil
}

func (db *DB) GetOriginalURLByShortURL(shortCode string) (*responseURL, error) {
	var result responseURL

	row := db.Conn.QueryRow("SELECT id, url, short_code, created_at, updated_at FROM urls WHERE short_code = ?", shortCode)
	if err := row.Scan(&result.ID, &result.URL, &result.ShortCode, &result.CreatedAt, &result.UpdatedAt); err != nil {
		return nil, fmt.Errorf("OriginalURLByShortURL %s: %w", shortCode, err)
	}

	return &result, nil
}
