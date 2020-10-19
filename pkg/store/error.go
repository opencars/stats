package store

import "errors"

var (
	// ErrRecordNotFound returned, when entity does not exist.
	ErrRecordNotFound = errors.New("record not found")
)
