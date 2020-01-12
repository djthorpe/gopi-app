package main

import (
	"context"
	"fmt"

	// Frameworks
	"github.com/djthorpe/gopi/v2"
)

////////////////////////////////////////////////////////////////////////////////

// List the EVENTS here you want to handle
var (
	EVENTS = []gopi.EventHandler{
		gopi.EventHandler{Name: "gopi.TimerEvent", Handler: TimerHandler},
	}
)

////////////////////////////////////////////////////////////////////////////////

// TimerHandler
func TimerHandler(_ context.Context, _ gopi.App, evt gopi.Event) {
	fmt.Println("Timer=", evt)
}
