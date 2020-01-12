package main

// This file is a list of UNITS you want to register
// as part of your application
import (
	_ "github.com/djthorpe/gopi/v2/unit/bus"
	_ "github.com/djthorpe/gopi/v2/unit/logger"
	_ "github.com/djthorpe/gopi/v2/unit/timer"
)

////////////////////////////////////////////////////////////////////////////////

// List the UNITS here you want to directly access
var (
	UNITS = []string{"timer"}
)
