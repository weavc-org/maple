package main

import (
	maple "github.com/weavc/maple/v1/pkg"
)

func NewApiEvent() *ApiEvent {
	return &ApiEvent{
		base: maple.NewBasicEvent(&maple.EventManifest{
			Namespace:   "maple.examples.api",
			Description: "Represents an event of an incoming Api event"}),
	}
}

type ApiEvent struct {
	maple.Event
	base *maple.BaseEvent
}

func (h *ApiEvent) Manifest() maple.EventManifest {
	return h.base.Manifest()
}

func (h *ApiEvent) Emit(v interface{}) {
	h.base.Emit(v)
}

func (h *ApiEvent) Register(handler maple.HandleEventFunc) error {
	return h.base.Register(handler)
}

func (h *ApiEvent) Walk(handler func(f maple.HandleEventFunc)) {
	h.base.Walk(handler)
}
