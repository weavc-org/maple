package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	eve "github.com/weavc/maple/v1/examples/http/event"
	"github.com/weavc/maple/v1/pkg"
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

	ev.Register(logRequestHandler, wrap.interceptRequestHandler)

	log.Fatal(http.ListenAndServe("0.0.0.0:3000", wrap))
}

type WrapEventMux struct {
	handler  http.Handler
	requests []*http.Request
}

func (l *WrapEventMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ev.Emit(&eve.HttpRequestEventArgs{R: r, W: w})
	l.handler.ServeHTTP(w, r)
}

func logRequestHandler(event pkg.Event, v interface{}) {
	go func() {
		args, _ := v.(*eve.HttpRequestEventArgs)
		log.Printf("%s %s %v", args.R.Method, args.R.URL.Path, time.Now())
	}()
}

func (l *WrapEventMux) interceptRequestHandler(event pkg.Event, v interface{}) {
	args, _ := v.(*eve.HttpRequestEventArgs)

	args.W.Header().Set("http-event-intercept", "true")
}
