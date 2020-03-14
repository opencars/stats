package model

import (
	"time"
)

type Authorization struct {
	ID        int64     `json:"-" db:"id"`
	Token     string    `json:"id" db:"token"`
	Name      *string   `json:"name,omitempty" db:"name"`
	Enabled   bool      `json:"enabled" db:"enabled"`
	Status    string    `json:"status" db:"status"`
	Error     *string   `json:"error,omitempty" db:"error"`
	IP        string    `json:"ip" db:"ip"`
	Time      time.Time `json:"timestamp" db:"timestamp"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
