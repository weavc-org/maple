package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/weavc/maple/v1/pkg"
)

var ev = NewApiEvent()

func main() {
	ev.Register(eventHandler)
	mux := http.NewServeMux()
	mux.HandleFunc("/api", helloworld)

	wrap := &WrapEventMux{mux}

	log.Fatal(http.ListenAndServe("0.0.0.0:3000", wrap))
}

type WrapEventMux struct {
	handler http.Handler
}

func (l *WrapEventMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ev.Emit(r)
	l.handler.ServeHTTP(w, r)
}

func helloworld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]string{"status": "200", "payload": "hello world"})
}

func eventHandler(event pkg.Event, v interface{}) {

	r, valid := v.(http.Request)
	if !valid {
		fmt.Printf("Invalid type on event handler. Expected http.Request")
	}

	log.Printf("%s %s %v", r.Method, r.URL.Path, time.Now())
}
