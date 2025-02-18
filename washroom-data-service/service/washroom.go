package service

import (
	"context"
	"washroom-data-service/models"
	"washroom-data-service/repository"
)

// WashroomService defines the interface for washroom-related business operations
type WashroomService interface {
	// Create adds a new washroom to the system
	Create(ctx context.Context, washroom *models.Washroom) error

	// GetByID retrieves a washroom by its unique identifier
	GetByID(ctx context.Context, id string) (*models.Washroom, error)

	// Update modifies an existing washroom's information
	Update(ctx context.Context, washroom *models.Washroom) error

	// Delete removes a washroom from the system
	Delete(ctx context.Context, id string) error

	// FindNearby locates washrooms within specified radius of coordinates
	FindNearby(ctx context.Context, lat, lng, radius float64) ([]models.Washroom, error)

	// FindInBuilding retrieves all washrooms in a specific building
	FindInBuilding(ctx context.Context, building string) ([]models.Washroom, error)

	// FindByFloor retrieves all washrooms on a specific floor of a building
	FindByFloor(ctx context.Context, building string, floor int) ([]models.Washroom, error)
}

// washroomService implements the WashroomService interface
type washroomService struct {
	repo         repository.WashroomRepository
	queryRepo    repository.LocationQueryRepository
	eventStore   repository.EventStore
	eventHandler repository.EventHandler
}

// NewWashroomService creates a new instance of WashroomService
func NewWashroomService(
	repo repository.WashroomRepository,
	queryRepo repository.LocationQueryRepository,
	eventStore repository.EventStore,
	eventHandler repository.EventHandler,
) WashroomService {
	return &washroomService{
		repo:         repo,
		queryRepo:    queryRepo,
		eventStore:   eventStore,
		eventHandler: eventHandler,
	}
}

// Create implements WashroomService.Create
func (s *washroomService) Create(ctx context.Context, washroom *models.Washroom) error {
	return s.repo.Create(ctx, washroom)
}

// GetByID implements WashroomService.GetByID
func (s *washroomService) GetByID(ctx context.Context, id string) (*models.Washroom, error) {
	return s.repo.GetByID(ctx, id)
}

// Update implements WashroomService.Update
func (s *washroomService) Update(ctx context.Context, washroom *models.Washroom) error {
	return s.repo.Update(ctx, washroom)
}

// Delete implements WashroomService.Delete
func (s *washroomService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

// FindNearby implements WashroomService.FindNearby
func (s *washroomService) FindNearby(ctx context.Context, lat, lng, radius float64) ([]models.Washroom, error) {
	return s.queryRepo.FindNearby(ctx, lat, lng, radius)
}

// FindInBuilding implements WashroomService.FindInBuilding
func (s *washroomService) FindInBuilding(ctx context.Context, building string) ([]models.Washroom, error) {
	return s.queryRepo.FindInBuilding(ctx, building)
}

// FindByFloor implements WashroomService.FindByFloor
func (s *washroomService) FindByFloor(ctx context.Context, building string, floor int) ([]models.Washroom, error) {
	return s.queryRepo.FindByFloor(ctx, building, floor)
}
