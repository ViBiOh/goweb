package main

import (
	"net/http"

	"github.com/ViBiOh/goweb/pkg/delay"
	"github.com/ViBiOh/goweb/pkg/dump"
	"github.com/ViBiOh/goweb/pkg/hello"
)

const (
	helloPath = "/hello/{name...}"
	dumpPath  = "/dump/"
	delayPath = "/delay/"
)

func newPort(config configuration, client client) http.Handler {
	mux := http.NewServeMux()

	mux.Handle(helloPath, hello.Handler(config.hello))
	mux.Handle(dumpPath, dump.Handler(client.telemetry.MeterProvider()))
	mux.Handle(delayPath, delay.Handler())

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTeapot)
	})

	return mux
}
