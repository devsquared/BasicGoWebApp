package basicgowebapp

import (
	"context"
	"html/template"
	"net/http"
	"sync"
)

// For lack of better name for this file, it is called app. There are many ways to organize your
// project structure but this template repo does not aim to be opinionated in that way. The main
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

// Handlers that have expensive operation - like loading templates from file and such - should be moved behind a
// sync to allow for it to be loaded once - the first time it is called.
func templateHandler(fileName string) http.Handler {
	var (
		init   sync.Once
		tpl    *template.Template
		tplerr error
	)
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			init.Do(func() {
				tpl, tplerr = template.ParseFiles(fileName) // here do the expensive pieces or just things that doesnt need to be done every load
			})
			if tplerr != nil {
				http.Error(w, tplerr.Error(), http.StatusInternalServerError)
				return
			}
			tpl.Execute(w, "data")
		},
	)
}

// reqRespHandler is an example of hiding the request and response in the handler. I highly recommend doing this
// for your handlers. The context of the request and response are tied to the handler so keep it close.
//
// NOTE: this kind of prep request and response in line can also help during testing. You are able to utilize tests as
// documentation even further with this approach as you show the reader what a request/response for this handler looks like.
func reqRespHandler() http.Handler {
	type request struct {
		ID string
	}

	type response struct {
		Payload string `json:"payload"`
	}

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			//...
		},
	)
}

func basicMiddleware(ctx context.Context, db DB, sc SomeClient) func(h http.Handler) http.Handler {
	// set up any dependencies needed for the set of handlers that this middleware will cover
	return func(h http.Handler) http.Handler {
		return h
	}
}

// These are just for example and basic typing for the methods above.
type DB string
type SomeClient string
