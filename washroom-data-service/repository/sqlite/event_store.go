package sqlite

import (
	"context"
	"database/sql"
	"encoding/json"

	"washroom-data-service/models"
	"washroom-data-service/repository"
)

type SQLiteEventStore struct {
	repository.BaseEventStore
	db *sql.DB
}

func NewSQLiteEventStore(db *sql.DB) *SQLiteEventStore {
	return &SQLiteEventStore{db: db}
}

func (s *SQLiteEventStore) SaveEvents(ctx context.Context, aggregateID string, events []models.Event) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, `
		INSERT INTO events (aggregate_id, type, data, version, timestamp)
		VALUES (?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, event := range events {
		data, err := event.ToJSON()
		if err != nil {
			return err
		}

		if _, err := stmt.ExecContext(ctx,
			event.GetAggregateID(),
			event.GetEventType(),
			data,
			event.GetVersion(),
			event.GetTimestamp()); err != nil {
			return err
		}

		if err := s.NotifyHandlers(event); err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (s *SQLiteEventStore) GetEvents(ctx context.Context, aggregateID string) ([]models.Event, error) {
	rows, err := s.db.QueryContext(ctx, `
		SELECT type, data, version, timestamp 
		FROM events 
		WHERE aggregate_id = ? 
		ORDER BY version ASC`, aggregateID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []models.Event
	for rows.Next() {
		var eventType string
		var data []byte
		var version int
		var timestamp string

		if err := rows.Scan(&eventType, &data, &version, &timestamp); err != nil {
			return nil, err
		}

		var event models.Event
		switch eventType {
		case "WashroomCreated":
			e := &models.WashroomCreatedEvent{}
			if err := json.Unmarshal(data, e); err != nil {
				return nil, err
			}
			event = e
		case "WashroomUpdated":
			e := &models.WashroomUpdatedEvent{}
			if err := json.Unmarshal(data, e); err != nil {
				return nil, err
			}
			event = e
		case "WashroomDeleted":
			e := &models.WashroomDeletedEvent{}
			if err := json.Unmarshal(data, e); err != nil {
				return nil, err
			}
			event = e
		}
		events = append(events, event)
	}

	if len(events) == 0 {
		return nil, repository.ErrAggregateNotFound
	}

	return events, rows.Err()
}

// ReconstructWashroom rebuilds a washroom's state by applying all events
func (s *SQLiteEventStore) ReconstructWashroom(ctx context.Context, aggregateID string) (*models.Washroom, error) {
	events, err := s.GetEvents(ctx, aggregateID)
	if err != nil {
		return nil, err
	}

	washroom := models.NewWashroom()
	for _, event := range events {
		if err := washroom.ApplyEvent(event); err != nil {
			return nil, err
		}
	}

	return washroom, nil
}
