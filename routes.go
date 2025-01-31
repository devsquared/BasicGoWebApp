package basicgowebapp

import (
	"context"
	"net/http"
)

// In this file, we list out the entire sitemap for the application. If your web application
// handles a route, it will be listed here. This really helps with troubleshooting and
// finding your way around the application.

// I like to keep this file strict to the routes that the app has. Handlers and middleware are to be created
// elsewhere. However for simple applications, you certainly can define and use those here.

func addRoutes(ctx context.Context, db DB, sc SomeClient, mux *http.ServeMux) {
	bMiddleware := basicMiddleware(ctx, db, sc)
	mux.Handle("GET /", bMiddleware(handleSomething(ctx)))
}
