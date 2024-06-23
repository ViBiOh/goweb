package main

import (
	"net/http"

	"github.com/ViBiOh/goweb/pkg/delay"
	"github.com/ViBiOh/goweb/pkg/dump"
	"github.com/ViBiOh/goweb/pkg/hello"
)

func newPort(config configuration, client clients) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /hello/{name...}", hello.Handler(config.hello))
	mux.Handle("/dump/", dump.Handler(client.telemetry.MeterProvider()))
	mux.HandleFunc("/delay/", delay.Handle)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTeapot)
	})

	return mux
}
