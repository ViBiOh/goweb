package main

import (
	"net/http"

	"github.com/ViBiOh/goweb/pkg/delay"
	"github.com/ViBiOh/goweb/pkg/dump"
	"github.com/ViBiOh/goweb/pkg/hello"
	"github.com/ViBiOh/httputils/v4/pkg/httputils"
)

func newPort(config configuration, clients clients, services services) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /hello/{name...}", hello.Handler(config.hello))
	mux.Handle("/dump/", dump.Handler(clients.telemetry.MeterProvider()))
	mux.HandleFunc("/delay/", delay.Handle)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTeapot)
	})

	return httputils.Handler(mux, clients.health,
		clients.telemetry.Middleware("http"),
		services.owasp.Middleware,
		services.cors.Middleware,
	)
}
