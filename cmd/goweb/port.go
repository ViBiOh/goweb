package main

import (
	"net/http"

	"github.com/ViBiOh/goweb/pkg/delay"
	"github.com/ViBiOh/goweb/pkg/dump"
	"github.com/ViBiOh/goweb/pkg/hello"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

const (
	helloPath = "/hello/"
	dumpPath  = "/dump/"
	delayPath = "/delay/"
)

func newPort(config configuration) http.Handler {
	mux := http.NewServeMux()

	mux.Handle(helloPath, otelhttp.WithRouteTag(helloPath, http.StripPrefix(helloPath, hello.Handler(config.hello))))
	mux.Handle(dumpPath, otelhttp.WithRouteTag(dumpPath, http.StripPrefix(dumpPath, dump.Handler())))
	mux.Handle(delayPath, otelhttp.WithRouteTag(delayPath, http.StripPrefix(delayPath, delay.Handler())))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTeapot)
	})

	return mux
}
