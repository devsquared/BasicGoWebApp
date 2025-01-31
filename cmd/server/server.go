package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
)

// This run pattern allows for us to decorate the main func a bit. The actual main then just becomes a simple
// method to call the run function and shut down the application in the case of an error. Utilize the run function
// as a way to call and do what is needed from the application, log, observe, shut down unnecessary things, and return
// a proper shut down error at the end.

// run allows for a great entry point to our application with a cancellable context. It also allows for us to
// easily decide on fatal or non-fatal error cases and handle accordingly. Further, this entry point for testing
// is perfect as you can run the application and cancel it (just a deferred cancel will do) at the end.
//
// NOTE: change up these args if needed
//   - examples of these could be os.stdin, config, or some func to gather input
func run(ctx context.Context, w io.Writer, args []string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	// do app stuff...

	return nil
}

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Stdout, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
