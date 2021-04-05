package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	eve "github.com/weavc/maple/examples/http/event"
	"github.com/weavc/maple/pkg"
)

var (
	ev = eve.NewHttpRequestEvent()
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]string{"status": "200", "payload": "hello world"})
	})

	wrap := &WrapEventMux{handler: mux, requests: make([]*http.Request, 0)}

	fmt.Printf("Registering handlers to event %s...", ev.GetManifest().Namespace)
	ev.Register(logRequestHandler, wrap.interceptRequestHandler)

	log.Fatal(http.ListenAndServe("0.0.0.0:3000", wrap))
}

type WrapEventMux struct {
	handler  http.Handler
	requests []*http.Request
}

func (l *WrapEventMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ev.Emit(&eve.HttpRequestEventArgs{Req: r, Res: w})
	l.handler.ServeHTTP(w, r)
}

func logRequestHandler(event pkg.Event, v interface{}) {
	go func() {
		args, _ := v.(*eve.HttpRequestEventArgs)
		log.Printf("%s %s %v", args.Req.Method, args.Req.URL.Path, time.Now())
	}()
}

func (l *WrapEventMux) interceptRequestHandler(event pkg.Event, v interface{}) {
	args, _ := v.(*eve.HttpRequestEventArgs)

	args.Res.Header().Set("http-event-intercept", "true")
}
