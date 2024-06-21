package main

import (
	"net/http"

	"github.com/ViBiOh/goweb/pkg/delay"
	"github.com/ViBiOh/goweb/pkg/dump"
	"github.com/ViBiOh/goweb/pkg/hello"
)

func newPort(config configuration, client client) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /hello/{name...}", hello.Handler(config.hello))
	mux.Handle("/dump/", dump.Handler(client.telemetry.MeterProvider()))
	mux.Handle("/delay/", delay.Handler())

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTeapot)
	})

	return mux
}
