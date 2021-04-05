package pkg

// Event
type Event interface {
	// Get the manifest of the event
	Manifest() EventManifest
	// Register event handlers with the event
	Register(handler ...HandleEventFunc) error
	// Emit will trigger any registered events
	Emit(v interface{})
}

// EventManifest is a breif overview of the event
type EventManifest struct {
	Namespace   string
	Description string
}

// HandleEventFunc defines an event handler
type HandleEventFunc func(event Event, v interface{})
