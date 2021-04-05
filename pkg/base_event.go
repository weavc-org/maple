package pkg

// Create & return a New BaseEvent with the provided namespace
// This can be extended or used as is. See examples/http for details
func NewBaseEvent(namespace string) *BaseEvent {
	return &BaseEvent{Manifest: &EventManifest{Namespace: namespace}}
}

// BaseEvent is an implementation of Event
// This can be used as is, or extended as shown in exmaples/http,
// using BaseEvent as a utility to support your event
type BaseEvent struct {
	Event

	Manifest   *EventManifest
	Registered []HandleEventFunc
}

func (h *BaseEvent) GetManifest() EventManifest {
	return *h.Manifest
}

func (h *BaseEvent) Emit(v interface{}) {
	for _, f := range h.Registered {
		f(h, v)
	}
}

func (h *BaseEvent) Register(handlers ...HandleEventFunc) error {
	h.Registered = append(h.Registered, handlers...)
	return nil
}

func (h *BaseEvent) Walk(handler func(f HandleEventFunc)) {
	for _, v := range h.Registered {
		handler(v)
	}
}
