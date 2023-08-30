package main

import (
	"net/http"
	"time"

	"github.com/alexedwards/flow"
)

func (app *application) server() *http.Server {
	return &http.Server{
		Addr:         ":10000",
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}

func (app *application) routes() http.Handler {
	mux := flow.New()

	mux.Use(app.recoverPanic)
	mux.Use(app.allowCors)

	mux.NotFound = http.HandlerFunc(notFound)
	mux.MethodNotAllowed = http.HandlerFunc(methodNotAllowed)

	mux.HandleFunc("/status", app.status, http.MethodGet)

	mux.Group(func(mux *flow.Mux) {
		mux.HandleFunc("/api/hello", app.handleHello, http.MethodGet)

	})

	return mux
}
