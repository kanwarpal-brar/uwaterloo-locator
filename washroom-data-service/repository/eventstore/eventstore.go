package eventstore

import (
	"context"
	"errors"

	"washroom-data-service/models"
)

var ErrAggregateNotFound = errors.New("aggregate not found")

type EventStore interface {
	SaveEvents(ctx context.Context, aggregateID string, events []models.Event) error
	GetEvents(ctx context.Context, aggregateID string) ([]models.Event, error)
}

type EventHandler interface {
	HandleEvent(event models.Event) error
}

type EventPublisher interface {
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
