package main

import (
	"context"
	"fmt"
	"os"
	"os/user"
	"time"

	// Frameworks
	"github.com/djthorpe/gopi/v2"
	"github.com/djthorpe/gopi/v2/app"
)

////////////////////////////////////////////////////////////////////////////////

func Main(app gopi.App, args []string) error {
	app.Timer().NewTicker(time.Second)

	// Print out the name
	fmt.Println("Hello, " + app.Flags().GetString("name", gopi.FLAG_NS_DEFAULT))

	// Wait for CTRL+C
	fmt.Println("Press CTRL+C to exit")
	app.WaitForSignal(context.Background(), os.Interrupt)

	// Return success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// BOOTSTRAP

func main() {
	if app, err := app.NewCommandLineTool(Main, EVENTS, UNITS...); err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else if user, err := user.Current(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		// Register -name flag
		app.Flags().FlagString("name", user.Name, "Name of user to print")

		// Run and exit
		os.Exit(app.Run())
	}
}
