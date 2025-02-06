package models

import (
	"time"
)

type Event interface {
	GetAggregateID() string
	GetEventType() string
	GetVersion() int
	GetTimestamp() time.Time
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
