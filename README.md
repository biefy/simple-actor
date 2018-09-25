# Simple Actor

This is a simple implementation of actor model in golang. It is mainly for simplify synchronization across go routines.

When multiple go routines access shared data, especially hierarchical data, it is pretty hard to sort things out. One 
approach to solve this problem is using an actor to process events.  This simple actor is an implementation with this idea.

## Simple to Use

```go

// Event definitions
const (
	EventAdd = iota
	EventMultiply
	EventSquare
)

// Create an actor
a := sa.New()

// Register event handlers
a.Register(EventAdd, func(args ...sa.Arg) {
	i1 := args[0].(*int)
	i2 := args[1].(int)

	*i1 += i2
})

a.Register(EventMultiply, func(args ...sa.Arg) {
	i1 := args[0].(*int)
	i2 := args[1].(int)

	*i1 *= i2
})

x := 2

// Cast events
a.Cast(EventAdd, &x, 1)         // x = 3
a.Cast(EventMultiply, &x, 10)	// x = 30

// Wait a second so that all events drain
<-time.After(time.Second)
fmt.Println("x =", x)

// Close the actor
a.Close()
```

For a complete example, see `examples/main.go`
