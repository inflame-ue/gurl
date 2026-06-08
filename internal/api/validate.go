package api

import (
	"errors"
	"fmt"
	"net/url"
)

func validateURL(rawURL string) error {
	if len(rawURL) == 0 {
		return errors.New("validation error: URL length cannot be zero")
	}

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	if parsedURL.Scheme == "" {
		return errors.New("validation error: URL scheme cannot be empty")
	}

	return nil
}
