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

type AuthStat struct {
	Token  string `json:"id" db:"token"`
	Name   string `json:"name,omitempty" db:"name"`
	Amount int64  `json:"amount" db:"amount"`
}

type StatsByIp struct {
	IP    string `json:"ip" db:"ip"`
	Total int64  `json:"total" db:"total"`
}

type TokenStat struct {
	Total   int64 `json:"total" db:"total"`
	Succeed int64 `json:"succeed" db:"succeed"`
	Failed  int64 `json:"failed" db:"failed"`
	// IPs     []StatsByIp `json:"ips" db:"ips"`
}
