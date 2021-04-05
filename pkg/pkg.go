package pkg

// Event
type Event interface {
	// Get the manifest of the event
	GetManifest() EventManifest
	// Register event handlers with the event
	Register(handler ...HandleEventFunc) error
	// Emit will trigger any registered events
	Emit(args interface{})
}

// EventManifest is a breif overview of the event
type EventManifest struct {
	Namespace   string
	Description string
	Data        map[string]interface{}
}

// HandleEventFunc defines an event handler
type HandleEventFunc func(event Event, args interface{})

// Create & return a New Event with the provided namespace
func NewEvent(namespace string) Event {
	return &BaseEvent{Manifest: &EventManifest{Namespace: namespace}}
}

// Create & return a New Event with the provided namespace
func NewEventFromManifest(manifest *EventManifest) Event {
	return &BaseEvent{Manifest: manifest}
}
