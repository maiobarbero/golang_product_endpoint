package main

import (
	"log"
	"net/http"
	"time"
)

type app struct {
	config config
}

type config struct {
	addr string // server address
}

func (app *app) mount() http.Handler {
	// For further improvement implement middlewares: auth, logger, rate limiting...

	mux := http.NewServeMux()

	// Routes registration

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("All good"))
	})

	return mux
}

func (app *app) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("server started at %s", app.config.addr)

	return srv.ListenAndServe()
}
