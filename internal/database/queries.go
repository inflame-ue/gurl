package database

import (
	"database/sql"
	"errors"
	"fmt"
)

type ShortURL struct {
	ID        int64  `json:"id"`
	URL       string `json:"url"`
	ShortCode string `json:"shortCode"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func (db *DB) CreateShortURL(longURL, shortCode string) (int64, error) {
	result, err := db.Conn.Exec("INSERT INTO urls (url, short_code) VALUES (?, ?)", longURL, shortCode)
	if err != nil {
		return 0, fmt.Errorf("CreateShortURL: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("CreateShortURL: %w", err)
	}
	return id, nil
}

func (db *DB) GetShortURLByID(id int64) (*ShortURL, error) {
	var result ShortURL

	row := db.Conn.QueryRow("SELECT id, url, short_code, created_at, updated_at FROM urls WHERE id = ?", id)
	if err := row.Scan(&result.ID, &result.URL, &result.ShortCode, &result.CreatedAt, &result.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("ShortURLsByID %d: no such url", id)
		}
		return nil, fmt.Errorf("ShortURLsByID %d: %w", id, err)
	}

	return &result, nil
}
