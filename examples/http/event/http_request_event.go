package event

import (
	"fmt"
	"net/http"

	maple "github.com/weavc/maple/pkg"
)

func NewHttpRequestEvent() *HttpRequestEvent {
	base := maple.NewBaseEvent("maple.examples.http")
	base.Manifest.Description = "Triggers on incoming HTTP Requests"

	return &HttpRequestEvent{
		base: base,
	}
}

type HttpRequestEvent struct {
	maple.Event
	base *maple.BaseEvent
}

type HttpRequestEventArgs struct {
	Res http.ResponseWriter
	Req *http.Request
}

func (h *HttpRequestEvent) GetManifest() maple.EventManifest {
	return h.base.GetManifest()
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
