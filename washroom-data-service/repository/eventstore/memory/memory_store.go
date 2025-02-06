package memory

import (
	"context"
	"sync"

	"washroom-data-service/models"
	"washroom-data-service/repository/eventstore"
)

type MemoryEventStore struct {
	eventstore.BaseEventStore
	mutex  sync.RWMutex
	events map[string][]models.Event
}

func NewMemoryEventStore() *MemoryEventStore {
	return &MemoryEventStore{
		events: make(map[string][]models.Event),
	}
}

func (s *MemoryEventStore) SaveEvents(ctx context.Context, aggregateID string, events []models.Event) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, event := range events {
		s.events[aggregateID] = append(s.events[aggregateID], event)
		if err := s.notifyHandlers(event); err != nil {
			return err
		}
	}
	return nil
}

func (s *MemoryEventStore) GetEvents(ctx context.Context, aggregateID string) ([]models.Event, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if events, exists := s.events[aggregateID]; exists {
		return events, nil
	}
	return nil, eventstore.ErrAggregateNotFound
}
