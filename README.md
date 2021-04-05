# maple

Very simple event library in go. 

- Define your event
  ```go
    manifest := maple.EventManifest{Namespace:"maple.events.foo"}
  ```
- Create a new Event 
  ```go
    event := NewBasicEvent(manifest)
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


