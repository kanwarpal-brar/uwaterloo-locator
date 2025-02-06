package models

type Aggregate interface {
	GetID() string
	GetVersion() int
	ApplyEvent(event Event) error
	GetUncommittedEvents() []Event
	ClearUncommittedEvents()
}

type BaseAggregate struct {
	ID                string
	Version           int
	UncommittedEvents []Event
}

func (a *BaseAggregate) GetID() string                 { return a.ID }
func (a *BaseAggregate) GetVersion() int               { return a.Version }
func (a *BaseAggregate) GetUncommittedEvents() []Event { return a.UncommittedEvents }
func (a *BaseAggregate) ClearUncommittedEvents()       { a.UncommittedEvents = nil }
