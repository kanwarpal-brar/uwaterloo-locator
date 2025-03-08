package models

import (
	"encoding/json"
	"time"
)

// Event represents a domain event that captures a state change in the system.
// Events are immutable and represent something that has happened in the past.
type Event interface {
	// GetAggregateID returns the ID of the aggregate this event belongs to
	GetAggregateID() string

	// GetEventType returns the type name of the event
	GetEventType() string

	// GetVersion returns the version number of the event in the aggregate's sequence
	GetVersion() int

	// GetTimestamp returns when the event occurred
	GetTimestamp() time.Time

	// ToJSON serializes the event to JSON for persistence
	ToJSON() ([]byte, error)
}

type BaseEvent struct {
	AggregateID string    `json:"aggregateId"`
	EventType   string    `json:"eventType"`
	Version     int       `json:"version"`
	Timestamp   time.Time `json:"timestamp"`
}

func (e BaseEvent) GetAggregateID() string  { return e.AggregateID }
func (e BaseEvent) GetEventType() string    { return e.EventType }
func (e BaseEvent) GetVersion() int         { return e.Version }
func (e BaseEvent) GetTimestamp() time.Time { return e.Timestamp }
func (e BaseEvent) ToJSON() ([]byte, error) { return json.Marshal(e) }

// Helper method for event serialization
func toJSON(e interface{}) ([]byte, error) {
	return json.Marshal(e)
}
