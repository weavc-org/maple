package event

import (
	"fmt"
	"net/http"

	maple "github.com/weavc/maple/v1/pkg"
)

func NewHttpRequestEvent() *HttpRequestEvent {
	return &HttpRequestEvent{
		base: maple.NewBasicEvent(&maple.EventManifest{
			Namespace:   "maple.examples.http",
			Description: "HTTP Request event"}),
	}
}

type HttpRequestEvent struct {
	maple.Event
	base *maple.BaseEvent
}

type HttpRequestEventArgs struct {
	W http.ResponseWriter
	R *http.Request
}

func (h *HttpRequestEvent) Manifest() maple.EventManifest {
	return h.base.Manifest()
}

func (h *HttpRequestEvent) Emit(v interface{}) {
	m, valid := v.(*HttpRequestEventArgs)
	if !valid {
		panic(fmt.Errorf("invalid args provided. Expected %s", "*ApiEventArgs"))
	}

	h.base.Emit(m)
}

func (h *HttpRequestEvent) Register(handlers ...maple.HandleEventFunc) error {
	return h.base.Register(handlers...)
}

func (h *HttpRequestEvent) Walk(handler func(f maple.HandleEventFunc)) {
	h.base.Walk(handler)
}
