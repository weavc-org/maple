package pkg

type Event interface {
	Manifest() EventManifest
	Register(handler HandleEventFunc) error
	Emit(v interface{})
}

type EventManifest struct {
	Namespace   string
	Description string
}

type HandleEventFunc func(event Event, v interface{})
