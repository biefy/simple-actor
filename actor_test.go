package simple_actor

import (
	"sync"
	"testing"
	"time"
)

func TestActor(t *testing.T) {
	const (
		Add Event = iota
		Multiply
	)

	start := 0
	lock := sync.RWMutex{}

	a := New()
	a.Register(Add, func(args ...Arg) {
		x := args[0].(*int)
		inc := args[1].(int)
		lock.Lock()
		*x += inc
		lock.Unlock()
	})

	a.Register(Multiply, func(args ...Arg) {
		x := args[0].(*int)
		mul := args[1].(int)
		lock.Lock()
		*x *= mul
		lock.Unlock()
	})

	a.Cast(Add, &start, 1)
	a.Cast(Multiply, &start, 3)

	if err := a.(*actor).waitForEmptyChan(time.Second * 5); err != nil {
		t.Errorf("failed to wait for channel drain: %v", err)
	}

	lock.RLock()
	if start != 3 {
		t.Errorf("start should be %d", 3)
	}
	lock.RUnlock()

	if err := a.Close(); err != nil {
		t.Errorf("failed to close actor: %v", err)
	}
}

func TestActor_Error(t *testing.T) {
	a := New()
	defer a.Close()

	if err := a.Cast(0); err == nil {
		t.Error("casting an unregistered event should fail")
	}

	if err := a.Register(0, nil); err == nil {
		t.Error("register an event with nil handler should fail")
	}

	if err := a.Register(0, func(args ...Arg) {}); err != nil {
		t.Errorf("register a valid event failed: %v", err)
	}

	if err := a.Register(0, func(args ...Arg) {}); err == nil {
		t.Errorf("re-register a valid event should fail")
	}
}
