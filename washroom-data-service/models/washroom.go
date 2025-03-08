package models

import (
	"time"
)

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Washroom struct {
	BaseAggregate
	Name            string    `json:"name"`
	Location        Location  `json:"location"`
	LocationUpdates int       `json:"locationUpdates"` // Track number of location updates
	Building        string    `json:"building"`
	Floor           int       `json:"floor"`
	Gender          string    `json:"gender"`
	IsAccessible    bool      `json:"isAccessible"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

// New constructor for creating a Washroom
func NewWashroom() *Washroom {
	return &Washroom{
		BaseAggregate: BaseAggregate{
			Version:           1,
			UncommittedEvents: make([]Event, 0),
		},
		LocationUpdates: 0,
	}
}

// ApplyEvent applies an event to the washroom aggregate
func (w *Washroom) ApplyEvent(event Event) error {
	switch e := event.(type) {
	case *WashroomCreatedEvent:
		w.applyWashroomCreated(e)
	case *WashroomUpdatedEvent:
		w.applyWashroomUpdated(e)
	case *WashroomDeletedEvent:
		w.applyWashroomDeleted(e)
	}

	w.Version = event.GetVersion()
	w.UncommittedEvents = append(w.UncommittedEvents, event)
	return nil
}

// applyWashroomCreated handles the WashroomCreated event
func (w *Washroom) applyWashroomCreated(e *WashroomCreatedEvent) {
	w.ID = e.AggregateID
	w.Name = e.Name
	w.Location = e.Location
	w.LocationUpdates = 1 // First location update
	w.Building = e.Building
	w.Floor = e.Floor
	w.Gender = e.Gender
	w.IsAccessible = e.IsAccessible
	w.CreatedAt = time.Now()
	w.UpdatedAt = time.Now()
}

// applyWashroomUpdated handles the WashroomUpdated event
func (w *Washroom) applyWashroomUpdated(e *WashroomUpdatedEvent) {
	if e.Name != "" {
		w.Name = e.Name
	}

	// Handle location update with moving average calculation
	if e.HasLocationUpdate {
		w.updateLocation(e.Location)
	}

	if e.Building != "" {
		w.Building = e.Building
	}
	if e.Floor != nil {
		w.Floor = *e.Floor
	}
	if e.Gender != "" {
		w.Gender = e.Gender
	}
	if e.IsAccessible != nil {
		w.IsAccessible = *e.IsAccessible
	}
	w.UpdatedAt = time.Now()
}

// applyWashroomDeleted handles the WashroomDeleted event
func (w *Washroom) applyWashroomDeleted(e *WashroomDeletedEvent) {
	// Mark as deleted or handle deletion logic
	// Currently this is a placeholder for future deletion logic
	// Could add a field like "IsDeleted" to the Washroom struct
}

// updateLocation calculates the moving average for the washroom location
func (w *Washroom) updateLocation(newLocation Location) {
	newLat := (w.Location.Latitude*float64(w.LocationUpdates) + newLocation.Latitude) / float64(w.LocationUpdates+1)
	newLng := (w.Location.Longitude*float64(w.LocationUpdates) + newLocation.Longitude) / float64(w.LocationUpdates+1)
	w.Location.Latitude = newLat
	w.Location.Longitude = newLng
	w.LocationUpdates++
}

// UpdateLocation updates the location and generates an appropriate event
func (w *Washroom) UpdateLocation(lat, lng float64) {
	updateEvent := &WashroomUpdatedEvent{
		BaseEvent: BaseEvent{
			AggregateID: w.ID,
			EventType:   "WashroomUpdated",
			Version:     w.Version + 1,
			Timestamp:   time.Now(),
		},
		HasLocationUpdate: true,
		Location: Location{
			Latitude:  lat,
			Longitude: lng,
		},
	}

	w.ApplyEvent(updateEvent)
}
