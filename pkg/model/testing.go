package model

import (
	"testing"
	"time"
)

func TestAuthorization(t *testing.T) *Authorization {
	t.Helper()

	name := "example"

	return &Authorization{
		Token:   "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		Name:    &name,
		Enabled: true,
		Status:  "succeed",
		Error:   nil,
		IP:      "127.0.0.1",
		Time:    time.Now().UTC(),
	}
}
