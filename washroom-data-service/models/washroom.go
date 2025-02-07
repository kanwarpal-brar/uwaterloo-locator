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
	Name         string    `json:"name"`
	Location     Location  `json:"location"`
	Building     string    `json:"building"`
	Floor        int       `json:"floor"`
	Gender       string    `json:"gender"`
	IsAccessible bool      `json:"isAccessible"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

func (w *Washroom) ApplyEvent(event Event) error {
	switch e := event.(type) {
	case *WashroomCreatedEvent:
		w.ID = e.AggregateID
		w.Name = e.Name
		w.Location = e.Location
		w.Building = e.Building
		w.Floor = e.Floor
		w.Gender = e.Gender
		w.IsAccessible = e.IsAccessible
	case *WashroomUpdatedEvent:
		// Apply updates based on event data
		if e.Name != "" {
			w.Name = e.Name
		}
		// ...handle other field updates...
	}
	w.Version++
	return nil
}
