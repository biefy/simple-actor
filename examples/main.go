package main

import (
	"fmt"
	"time"

	sa "github.com/biefy/simple-actor"
)

const (
	EventAdd = iota
	EventMultiply
	EventSquare
)

func main() {
	a := sa.New()
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

	a.Register(EventSquare, func(args ...sa.Arg) {
		i1 := args[0].(*int)
		*i1 *= *i1
	})

	x := 2
	a.Cast(EventAdd, &x, 1)  		// x = 3
	a.Cast(EventMultiply, &x, 10)	// x = 30
	a.Cast(EventAdd, &x, 100)		// x = 130
	a.Cast(EventMultiply, &x, 3)	// x = 390
	a.Cast(EventSquare, &x)			// x = 152100

	<-time.After(time.Second)
	fmt.Println("x =", x)
}