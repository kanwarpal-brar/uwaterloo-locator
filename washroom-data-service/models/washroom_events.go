package models

import (
	"encoding/json"
)

// WashroomCreatedEvent represents the creation of a washroom
type WashroomCreatedEvent struct {
	BaseEvent
	Name         string   `json:"name"`
	Location     Location `json:"location"`
	Building     string   `json:"building"`
	Floor        int      `json:"floor"`
	Gender       string   `json:"gender"`
	IsAccessible bool     `json:"isAccessible"`
}

// ToJSON serializes the WashroomCreatedEvent to JSON
func (e WashroomCreatedEvent) ToJSON() ([]byte, error) {
	return json.Marshal(e)
}

// WashroomUpdatedEvent represents an update to a washroom's details
type WashroomUpdatedEvent struct {
	BaseEvent
	Name              string   `json:"name,omitempty"`
	HasLocationUpdate bool     `json:"hasLocationUpdate"`
	Location          Location `json:"location,omitempty"`
	Building          string   `json:"building,omitempty"`
	Floor             *int     `json:"floor,omitempty"`
	Gender            string   `json:"gender,omitempty"`
	IsAccessible      *bool    `json:"isAccessible,omitempty"`
}

// ToJSON serializes the WashroomUpdatedEvent to JSON
func (e WashroomUpdatedEvent) ToJSON() ([]byte, error) {
	return json.Marshal(e)
}

// WashroomDeletedEvent represents the deletion of a washroom
type WashroomDeletedEvent struct {
	BaseEvent
}

// ToJSON serializes the WashroomDeletedEvent to JSON
func (e WashroomDeletedEvent) ToJSON() ([]byte, error) {
	return json.Marshal(e)
}

// CreateWashroomEvent creates a new WashroomCreatedEvent
func CreateWashroomEvent(washroom *Washroom) *WashroomCreatedEvent {
	return &WashroomCreatedEvent{
		BaseEvent: BaseEvent{
			AggregateID: washroom.ID,
			EventType:   "WashroomCreated",
			Version:     1,
			Timestamp:   washroom.CreatedAt,
		},
		Name:         washroom.Name,
		Location:     washroom.Location,
		Building:     washroom.Building,
		Floor:        washroom.Floor,
		Gender:       washroom.Gender,
		IsAccessible: washroom.IsAccessible,
	}
}
