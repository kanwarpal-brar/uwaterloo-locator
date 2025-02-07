package repository

import (
	"context"

	"washroom-data-service/models"
)

// WashroomRepository defines the interface for persisting and retrieving washroom data.
// It follows the repository pattern to abstract the underlying storage mechanism.
type WashroomRepository interface {
	// Create persists a new washroom and assigns it a unique ID
	Create(ctx context.Context, washroom *models.Washroom) error

	// GetByID retrieves a washroom by its unique identifier
	GetByID(ctx context.Context, id string) (*models.Washroom, error)

	// Update modifies an existing washroom's data
	Update(ctx context.Context, washroom *models.Washroom) error

	// Delete removes a washroom from the system
	Delete(ctx context.Context, id string) error
}

// LocationQueryRepository defines specialized queries for finding washrooms by location.
// This interface implements the CQRS pattern by separating read operations
// from write operations (handled by WashroomRepository).
type LocationQueryRepository interface {
	// FindNearby returns washrooms within radius (meters) of the given location
	FindNearby(ctx context.Context, lat, lng, radius float64) ([]models.Washroom, error)
	// FindInBuilding returns all washrooms in a specific building
	FindInBuilding(ctx context.Context, building string) ([]models.Washroom, error)
	// FindByFloor returns all washrooms on a specific floor of a building
	FindByFloor(ctx context.Context, building string, floor int) ([]models.Washroom, error)
}
