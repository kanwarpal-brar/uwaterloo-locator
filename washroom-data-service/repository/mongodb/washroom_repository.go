package mongodb

import (
	"context"

	"washroom-data-service/models"
	"washroom-data-service/models/mongodb"
	"washroom-data-service/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(db *mongo.Database) repository.WashroomRepository {
	return &mongoRepository{
		collection: db.Collection("washrooms"),
	}
}

func (r *mongoRepository) Create(ctx context.Context, w *models.Washroom) error {
	doc := mongodb.ToDocument(w)
	result, err := r.collection.InsertOne(ctx, doc)
	if err != nil {
		return err
	}
	w.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return nil
}

func (r *mongoRepository) GetByID(ctx context.Context, id string) (*models.Washroom, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, repository.ErrNotFound
	}

	var doc mongodb.WashroomDocument
	if err := r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&doc); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return mongodb.FromDocument(&doc), nil
}

func (r *mongoRepository) Update(ctx context.Context, w *models.Washroom) error {
	objectID, err := primitive.ObjectIDFromHex(w.ID)
	if err != nil {
		return repository.ErrNotFound
	}

	doc := mongodb.ToDocument(w)
	result, err := r.collection.ReplaceOne(ctx, bson.M{"_id": objectID}, doc)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return repository.ErrNotFound
	}
	return nil
}

func (r *mongoRepository) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return repository.ErrNotFound
	}

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return repository.ErrNotFound
	}
	return nil
}
