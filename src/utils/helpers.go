package coreengine

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

// HTTPError represents an error that occurred during an HTTP request
type HTTPError struct {
	Code    int
	Message string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTP error %d: %s", e.Code, e.Message)
}

// getHTTPClient returns a new HTTP client with a 10-second timeout
func getHTTPClient() *http.Client {
	return &http.Client{
		Timeout: 10 * time.Second,
	}
}

// parseInt parses a string into an integer
func parseInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, errors.New("invalid integer")
	}
	return i, nil
}

// parseBoolean parses a string into a boolean
func parseBoolean(s string) (bool, error) {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return false, errors.New("invalid boolean")
	}
	return b, nil
}

// logError logs an error with a message
func logError(msg string, err error) {
	log.Printf("%s: %v", msg, err)
}
func getNow() time.Time {
	return time.Now().UTC()
}
func getToday() time.Time {
	return getNow().Truncate(24 * time.Hour)
}