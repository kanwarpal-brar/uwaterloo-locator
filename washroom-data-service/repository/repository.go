package repository

import (
	"context"

	"washroom-data-service/models"
)

// WashroomRepository defines operations for managing washroom data
type WashroomRepository interface {
	Create(ctx context.Context, washroom *models.Washroom) error
	GetByID(ctx context.Context, id string) (*models.Washroom, error)
	Update(ctx context.Context, washroom *models.Washroom) error
	Delete(ctx context.Context, id string) error
}

// LocationQueryRepository defines operations for querying location data
type LocationQueryRepository interface {
	// FindNearby returns washrooms within radius (meters) of the given location
	FindNearby(ctx context.Context, lat, lng, radius float64) ([]models.Washroom, error)
	// FindInBuilding returns all washrooms in a specific building
	FindInBuilding(ctx context.Context, building string) ([]models.Washroom, error)
	// FindByFloor returns all washrooms on a specific floor of a building
	FindByFloor(ctx context.Context, building string, floor int) ([]models.Washroom, error)
}
