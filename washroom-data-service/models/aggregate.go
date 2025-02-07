package models

// Aggregate defines the interface for domain aggregates that can handle events
// and maintain their state through event sourcing.
type Aggregate interface {
	// GetID returns the unique identifier of the aggregate
	GetID() string

	// GetVersion returns the current version of the aggregate
	GetVersion() int

	// ApplyEvent applies an event to the aggregate, updating its state
	ApplyEvent(event Event) error

	// GetUncommittedEvents returns all new events that haven't been persisted
	GetUncommittedEvents() []Event

	// ClearUncommittedEvents removes all uncommitted events after they've been persisted
	ClearUncommittedEvents()
}

// BaseAggregate provides a basic implementation of the Aggregate interface
type BaseAggregate struct {
	// ID uniquely identifies the aggregate
	ID string

	// Version represents the current version of the aggregate
	Version int

	// UncommittedEvents contains events that haven't been persisted
	UncommittedEvents []Event
}

// GetID returns the aggregate's unique identifier
func (a *BaseAggregate) GetID() string { return a.ID }

// GetVersion returns the aggregate's current version
func (a *BaseAggregate) GetVersion() int { return a.Version }

// GetUncommittedEvents returns all new events that haven't been persisted
func (a *BaseAggregate) GetUncommittedEvents() []Event { return a.UncommittedEvents }

// ClearUncommittedEvents removes all uncommitted events after persistence
func (a *BaseAggregate) ClearUncommittedEvents() { a.UncommittedEvents = nil }
