package mongodb

import (
	"time"

	"washroom-data-service/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Location struct {
	Type        string    `bson:"type"`
	Coordinates []float64 `bson:"coordinates"`
}

type WashroomDocument struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	Location     Location           `bson:"location"`
	Building     string             `bson:"building"`
	Floor        int                `bson:"floor"`
	Gender       string             `bson:"gender"`
	IsAccessible bool               `bson:"isAccessible"`
	Version      int                `bson:"version"`
	CreatedAt    time.Time          `bson:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt"`
}

func ToDocument(w *models.Washroom) *WashroomDocument {
	return &WashroomDocument{
		Name: w.Name,
		Location: Location{
			Type:        "Point",
			Coordinates: []float64{w.Location.Longitude, w.Location.Latitude},
		},
		Building:     w.Building,
		Floor:        w.Floor,
		Gender:       w.Gender,
		IsAccessible: w.IsAccessible,
		Version:      w.Version,
		CreatedAt:    w.CreatedAt,
		UpdatedAt:    w.UpdatedAt,
	}
}

func FromDocument(doc *WashroomDocument) *models.Washroom {
	return &models.Washroom{
		BaseAggregate: models.BaseAggregate{
			ID:      doc.ID.Hex(),
			Version: doc.Version,
		},
		Name: doc.Name,
		Location: models.Location{
			Latitude:  doc.Location.Coordinates[1],
			Longitude: doc.Location.Coordinates[0],
		},
		Building:     doc.Building,
		Floor:        doc.Floor,
		Gender:       doc.Gender,
		IsAccessible: doc.IsAccessible,
		CreatedAt:    doc.CreatedAt,
		UpdatedAt:    doc.UpdatedAt,
	}
}
