package pkg

func NewBasicEvent(manifest *EventManifest) *BaseEvent {
	return &BaseEvent{manifest: manifest}
}

type BaseEvent struct {
	Event

	manifest   *EventManifest
	registered []HandleEventFunc
}

// Emit an event
func (h *BaseEvent) Manifest() EventManifest {
	return *h.manifest
}

// Emit an event
func (h *BaseEvent) Emit(v interface{}) {
	for _, f := range h.registered {
		go f(h, v)
	}
}

// Register a function to an event
// This will also register the event if it is not found in the event map
func (h *BaseEvent) Register(handler HandleEventFunc) error {

	h.registered = append(h.registered, handler)

	return nil
}

// Take walk through the registered handlers
func (h *BaseEvent) Walk(handler func(f HandleEventFunc)) {
	for _, v := range h.registered {
		handler(v)
	}
}
