package model

// EventKind ...
type EventKind string

const (
	// EventAuthorizationKind ...
	EventAuthorizationKind EventKind = "authorization"
)

// Event has a pre-defined payload.
type Event struct {
	Kind EventKind   `json:"kind"`
	Data interface{} `json:"data"`
}
