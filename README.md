# maple

Very simple event library in go. 

[![Go](https://github.com/weavc/maple/actions/workflows/go_build_and_test.yml/badge.svg)](https://github.com/weavc/maple/actions/workflows/go_build_and_test.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/weavc/maple.svg)](https://pkg.go.dev/github.com/weavc/maple)

### Usage
- Import the library
  ```go
    import (
      maple "github.com/weavc/maple/pkg"
    )
  ```
- Create a new Event 
  ```go
    event := maple.NewEvent("maple.events.foo")
  ```
- Register your handlers
  ```go
    event.Register(func(ev maple.Event, v interface{}) {
      // do something
    })
  ```
- Start emitting events
  ```go
    event.Emit(payload)
  ```

-----

### Handling events
When you register functions to an event, these must implement our `HandleEventFunc` type.
```go
type HandleEventFunc func(event Event, args interface{})
```

When registered to an event, the function is added to a stack of methods that will be called everytime the event is triggered. When called it will be passed the `Event` structure/interface calling it and the value that is provided when the event is triggered.

The library leaves how the event is handled to the handlers themselves. This allows the handler to decide how it wants to handle the event. i.e. 
- Spawn a go routine
```go
func eventHandler(event pkg.Event, v interface{}) {
  go func() {
    // do stuff
  }()
}
```
- Type the args
```go
func eventHandler(event pkg.Event, v interface{}) {
  s, valid := v.(string)
  if !valid {
    panic(fmt.Errorf("invalid args provided. Expected %s", "*ApiEventArgs"))
  }
}
```

