package mongodb

import (
	"context"

	"washroom-data-service/models"
	"washroom-data-service/models/mongodb"
	"washroom-data-service/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoLocationQuery struct {
	collection *mongo.Collection
}

func NewMongoLocationQuery(db *mongo.Database) repository.LocationQueryRepository {
	return &mongoLocationQuery{
		collection: db.Collection("washrooms"),
	}
}

func (r *mongoLocationQuery) FindNearby(ctx context.Context, lat, lng, radius float64) ([]models.Washroom, error) {
	filter := bson.M{
		"location": bson.M{
			"$near": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{lng, lat},
				},
				"$maxDistance": radius,
			},
		},
	}
	return r.find(ctx, filter)
}

func (r *mongoLocationQuery) FindInBuilding(ctx context.Context, building string) ([]models.Washroom, error) {
	filter := bson.M{"building": building}
	return r.find(ctx, filter)
}

func (r *mongoLocationQuery) FindByFloor(ctx context.Context, building string, floor int) ([]models.Washroom, error) {
	filter := bson.M{
		"building": building,
		"floor":    floor,
	}
	return r.find(ctx, filter)
}

func (r *mongoLocationQuery) find(ctx context.Context, filter interface{}) ([]models.Washroom, error) {
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var docs []mongodb.WashroomDocument
	if err := cursor.All(ctx, &docs); err != nil {
		return nil, err
	}

	washrooms := make([]models.Washroom, len(docs))
	for i, doc := range docs {
		washrooms[i] = *mongodb.FromDocument(&doc)
	}
	return washrooms, nil
}
