package pkg

func NewBasicEvent(manifest *EventManifest) *BaseEvent {
	return &BaseEvent{manifest: manifest}
}

type BaseEvent struct {
	Event

	manifest   *EventManifest
	registered []HandleEventFunc
}

// Return event manifest
func (h *BaseEvent) Manifest() EventManifest {
	return *h.manifest
}

// Emit an event
func (h *BaseEvent) Emit(v interface{}) {
	for _, f := range h.registered {
		f(h, v)
	}
}

// Register handlers to an event
func (h *BaseEvent) Register(handlers ...HandleEventFunc) error {
	h.registered = append(h.registered, handlers...)
	return nil
}

// Take walk through the registered event handlers
func (h *BaseEvent) Walk(handler func(f HandleEventFunc)) {
	for _, v := range h.registered {
		handler(v)
	}
}
