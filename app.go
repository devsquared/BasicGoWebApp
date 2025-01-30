package basicgowebapp

import (
	"context"
	"net/http"
)

// For lack of better name for this file, it is called app. There are many ways to organize your
// project structure but this template does not aim to be opinionated in that way. The main
// point of this file is to show the structure of handlers. Just keep in mind, that you can
// separate your handlers in whatever manner you want to to organize. Handlers are the interface
// the web will use to communicate to your business logic so you can organize your logic beneath
// the handlers as well.

// handleSomething is a func that returns a handler. This allows for a basic closure structure around
// your handlers. This allows for you to set things up that may be needed for the handler to do its work.
func handleSomething(ctx context.Context) http.Handler {
	// Example: thing := prepareThing()
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// use thing to handle request
			// Example of utilizing logger from context: ctx.Logger().Info(r.Context(), "msg", "handleSomething")

			encode(w, r, http.StatusOK, "success")
		},
	)
}
