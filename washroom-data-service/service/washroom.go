package service

import (
	"context"
	"time"
	"washroom-data-service/models"
	"washroom-data-service/repository"

	"github.com/google/uuid"
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
	// Create a unique ID for the washroom if not provided
	if washroom.ID == "" {
		washroom.ID = uuid.New().String()
	}

	// Create the initial event
	createdEvent := &models.WashroomCreatedEvent{
		BaseEvent: models.BaseEvent{
			AggregateID: washroom.ID,
			EventType:   "WashroomCreated",
			Version:     1,
			Timestamp:   time.Now(),
		},
		Name:         washroom.Name,
		Location:     washroom.Location,
		Building:     washroom.Building,
		Floor:        washroom.Floor,
		Gender:       washroom.Gender,
		IsAccessible: washroom.IsAccessible,
	}

	// Apply the event to update the washroom's state
	if err := washroom.ApplyEvent(createdEvent); err != nil {
		return err
	}

	// Save the event
	return s.eventStore.SaveEvents(ctx, washroom.ID, []models.Event{createdEvent})
}

// GetByID implements WashroomService.GetByID
func (s *washroomService) GetByID(ctx context.Context, id string) (*models.Washroom, error) {
	// For event sourcing, reconstruct the washroom from events
	events, err := s.eventStore.GetEvents(ctx, id)
	if err != nil {
		return nil, err
	}

	// Reconstruct the washroom by replaying events
	washroom := models.NewWashroom()
	for _, event := range events {
		if err := washroom.ApplyEvent(event); err != nil {
			return nil, err
		}
	}

	return washroom, nil
}

// Update implements WashroomService.Update
func (s *washroomService) Update(ctx context.Context, washroom *models.Washroom) error {
	// First, get the current state by reconstructing from events
	currentWashroom, err := s.GetByID(ctx, washroom.ID)
	if err != nil {
		return err
	}

	// Create the update event
	updateEvent := &models.WashroomUpdatedEvent{
		BaseEvent: models.BaseEvent{
			AggregateID: washroom.ID,
			EventType:   "WashroomUpdated",
			Version:     currentWashroom.Version + 1,
			Timestamp:   time.Now(),
		},
		Name: washroom.Name,
	}

	// Check if location is being updated
	if washroom.Location.Latitude != currentWashroom.Location.Latitude ||
		washroom.Location.Longitude != currentWashroom.Location.Longitude {
		updateEvent.HasLocationUpdate = true
		updateEvent.Location = washroom.Location
	}

	// Set other fields if they've changed
	if washroom.Building != currentWashroom.Building {
		updateEvent.Building = washroom.Building
	}
	if washroom.Floor != currentWashroom.Floor {
		floor := washroom.Floor
		updateEvent.Floor = &floor
	}
	if washroom.Gender != currentWashroom.Gender {
		updateEvent.Gender = washroom.Gender
	}
	if washroom.IsAccessible != currentWashroom.IsAccessible {
		accessible := washroom.IsAccessible
		updateEvent.IsAccessible = &accessible
	}

	// Apply the event to update the current washroom's state
	if err := currentWashroom.ApplyEvent(updateEvent); err != nil {
		return err
	}

	// Copy updated state back to the input washroom
	*washroom = *currentWashroom

	// Save the event
	return s.eventStore.SaveEvents(ctx, washroom.ID, []models.Event{updateEvent})
}

// Delete implements WashroomService.Delete
func (s *washroomService) Delete(ctx context.Context, id string) error {
	// Get current washroom to check it exists and get version
	currentWashroom, err := s.GetByID(ctx, id)
	if err != nil {
		return err
	}

	// Create delete event
	deleteEvent := &models.WashroomDeletedEvent{
		BaseEvent: models.BaseEvent{
			AggregateID: id,
			EventType:   "WashroomDeleted",
			Version:     currentWashroom.Version + 1,
			Timestamp:   time.Now(),
		},
	}

	// Apply the event
	if err := currentWashroom.ApplyEvent(deleteEvent); err != nil {
		return err
	}

	// Save the event
	return s.eventStore.SaveEvents(ctx, id, []models.Event{deleteEvent})
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
