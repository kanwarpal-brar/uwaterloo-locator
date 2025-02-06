package models

type WashroomCreatedEvent struct {
	BaseEvent
	Name         string   `json:"name"`
	Location     Location `json:"location"`
	Building     string   `json:"building"`
	Floor        int      `json:"floor"`
	Gender       string   `json:"gender"`
	IsAccessible bool     `json:"isAccessible"`
}

type WashroomUpdatedEvent struct {
	BaseEvent
	Name         string   `json:"name,omitempty"`
	Location     Location `json:"location,omitempty"`
	Building     string   `json:"building,omitempty"`
	Floor        *int     `json:"floor,omitempty"`
	Gender       string   `json:"gender,omitempty"`
	IsAccessible *bool    `json:"isAccessible,omitempty"`
}

type WashroomDeletedEvent struct {
	BaseEvent
}
