# maple

Very simple event library in go. 
- Import the library
  ```go
    import (
      maple "github.com/weavc/maple/v1/pkg"
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


