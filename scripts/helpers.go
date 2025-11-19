package helpers

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"strings"
	"time"
)

// GenerateRandomString generates a random string of the given length.
func GenerateRandomString(length int) (string, error) {
	if length <= 0 {
		return "", errors.New("length must be greater than 0")
	}

	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes)[:length], nil
}

// FormatTimestamp formats a timestamp into a human-readable string.
func FormatTimestamp(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// ValidateEmail checks if the email address is valid.
func ValidateEmail(email string) bool {
	if strings.Count(email, "@") != 1 {
		return false
	}

	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}

	if parts[0] == "" || parts[1] == "" {
		return false
	}

	return true
}