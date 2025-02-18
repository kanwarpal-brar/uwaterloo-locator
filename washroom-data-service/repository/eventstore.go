package repository

import (
	"context"
	"errors"

	"washroom-data-service/models"
)

var ErrAggregateNotFound = errors.New("aggregate not found")

// EventStore defines the interface for an event sourcing storage system
type EventStore interface {
	// SaveEvents atomically persists new events for an aggregate
	// Returns an error if the save operation fails
	SaveEvents(ctx context.Context, aggregateID string, events []models.Event) error

	// GetEvents retrieves all events for an aggregate in chronological order
	// Returns ErrAggregateNotFound if the aggregate doesn't exist
	GetEvents(ctx context.Context, aggregateID string) ([]models.Event, error)
}

// EventHandler processes events after they are persisted
type EventHandler interface {
	// HandleEvent processes a single event
	// Returns an error if the handling fails
	HandleEvent(event models.Event) error
}

// EventPublisher distributes events to external systems
type EventPublisher interface {
	// PublishEvent sends an event to external subscribers
	// Returns an error if publishing fails
	PublishEvent(event models.Event) error
}

// BaseEventStore provides common event store functionality
type BaseEventStore struct {
	handlers  []EventHandler
	publisher EventPublisher
}

func (s *BaseEventStore) AddHandler(handler EventHandler) {
	s.handlers = append(s.handlers, handler)
}

// NotifyHandlers notifies all registered handlers about an event
func (s *BaseEventStore) NotifyHandlers(event models.Event) error {
	for _, handler := range s.handlers {
		if err := handler.HandleEvent(event); err != nil {
			return err
		}
	}
	if s.publisher != nil {
		return s.publisher.PublishEvent(event)
	}
	return nil
}
