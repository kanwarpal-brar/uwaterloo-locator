package mongodb

import (
	"context"
	"time"

	"washroom-data-service/models"
	"washroom-data-service/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type EventDocument struct {
	AggregateID string    `bson:"aggregateId"`
	Type        string    `bson:"type"`
	Data        []byte    `bson:"data"`
	Version     int       `bson:"version"`
	Timestamp   time.Time `bson:"timestamp"`
}

type MongoEventStore struct {
	repository.BaseEventStore
	collection *mongo.Collection
}

func NewMongoEventStore(db *mongo.Database) *MongoEventStore {
	return &MongoEventStore{
		collection: db.Collection("events"),
	}
}

func (s *MongoEventStore) SaveEvents(ctx context.Context, aggregateID string, events []models.Event) error {
	// TODO: Implementation for saving events to MongoDB
	return nil
}

func (s *MongoEventStore) GetEvents(ctx context.Context, aggregateID string) ([]models.Event, error) {
	cursor, err := s.collection.Find(ctx, bson.M{"aggregateId": aggregateID})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, repository.ErrAggregateNotFound
		}
		return nil, err
	}
	defer cursor.Close(ctx)

	var events []models.Event
	// TODO: Implementation for converting documents to events
	return events, nil
}
