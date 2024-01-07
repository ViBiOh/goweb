package main

import (
	"net/http"

	"github.com/ViBiOh/goweb/pkg/delay"
	"github.com/ViBiOh/goweb/pkg/dump"
	"github.com/ViBiOh/goweb/pkg/hello"
)

const (
	helloPath = "/hello/"
	dumpPath  = "/dump/"
	delayPath = "/delay/"
)

func newPort(config configuration) http.Handler {
	mux := http.NewServeMux()

	mux.Handle(helloPath, http.StripPrefix(helloPath, hello.Handler(config.hello)))
	mux.Handle(dumpPath, http.StripPrefix(dumpPath, dump.Handler()))
	mux.Handle(delayPath, http.StripPrefix(delayPath, delay.Handler()))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTeapot)
	})

	return mux
}
