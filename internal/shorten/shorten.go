package shorten

import "crypto/rand"

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func ShortenURL(length int) string {
	randomBytes := make([]byte, length)
	rand.Read(randomBytes)

	result := make([]byte, length)
	for i := range result {
		randomIndex := int(randomBytes[i]) % len(charset)
		result[i] = charset[randomIndex]
	}

	return string(result)
}
