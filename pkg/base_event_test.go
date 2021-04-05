package pkg

import "testing"

func TestBaseEventManifest(t *testing.T) {
	name := "maple.test.base_event_manifest"
	desc := "Testing event for maple"
	var data map[string]interface{} = map[string]interface{}{"test": "foo"}

	ev := NewBaseEvent(name)

	if ev.Manifest.Namespace != name {
		t.Errorf("namespace doesn't match provided name. expected: %s, got: %s",
			name,
			ev.Manifest.Namespace)
	}

	ev.Manifest.Description = desc
	ev.Manifest.Data = data

	manifest := ev.GetManifest()
	if manifest.Description != desc {
		t.Errorf("Description doesn't match. expected: %s, got: %s",
			desc,
			manifest.Description)
	}
}

func TestBaseEventRegister(t *testing.T) {
	name := "maple.test.base_event_register"

	ev := NewBaseEvent(name)

	ev.Register(func(event Event, v interface{}) {})

	if len(ev.Registered) != 1 {
		t.Errorf("Registered length is incorrect. expected: 1, got: %v", len(ev.Registered))
	}

	ev.Register(func(event Event, v interface{}) {}, func(event Event, v interface{}) {})

	if len(ev.Registered) != 3 {
		t.Errorf("Registered length is incorrect. expected: 3, got: %v", len(ev.Registered))
	}

	var i int = 0
	ev.Walk(func(f HandleEventFunc) {
		i = i + 1
	})

	if i != 3 {
		t.Errorf("Value i is incorrect. expected: 3, got: %v", i)
	}

}

func TestBaseEventEmit(t *testing.T) {
	name := "maple.test.base_event_emit"

	ev := NewBaseEvent(name)
	var hit1 bool = false

	ev.Register(func(event Event, v interface{}) {
		hit1 = v.(bool)
	})

	ev.Emit(true)

	if !hit1 {
		t.Errorf("Value of hit1 was incorrect. expected: true, got: %v", hit1)
	}

	hit1 = false
	var hit2 bool = false
	var hit3 bool = false

	ev.Register(func(event Event, args interface{}) {
		hit2 = true
	}, func(event Event, args interface{}) {
		hit3 = true
	})

	ev.Emit(true)

	if !hit1 || !hit2 || !hit3 {
		t.Errorf("Value of hits were incorrect. expected: true, true, true, got: %v, %v, %v", hit1, hit2, hit3)
	}

}
