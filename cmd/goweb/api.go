package main

import (
	"flag"
	"net/http"
	"os"
	"strings"

	"github.com/ViBiOh/goweb/pkg/dump"
	"github.com/ViBiOh/goweb/pkg/hello"
	"github.com/ViBiOh/httputils/v3/pkg/alcotest"
	"github.com/ViBiOh/httputils/v3/pkg/cors"
	"github.com/ViBiOh/httputils/v3/pkg/httputils"
	"github.com/ViBiOh/httputils/v3/pkg/logger"
	"github.com/ViBiOh/httputils/v3/pkg/owasp"
	"github.com/ViBiOh/httputils/v3/pkg/prometheus"
)

const (
	helloPath = "/hello"
	dumpPath  = "/dump"
)

func main() {
	fs := flag.NewFlagSet("api", flag.ExitOnError)

	serverConfig := httputils.Flags(fs, "")
	alcotestConfig := alcotest.Flags(fs, "")
	loggerConfig := logger.Flags(fs, "logger")
	prometheusConfig := prometheus.Flags(fs, "prometheus")
	owaspConfig := owasp.Flags(fs, "")
	corsConfig := cors.Flags(fs, "cors")

	helloConfig := hello.Flags(fs, "")

	logger.Fatal(fs.Parse(os.Args[1:]))

	alcotest.DoAndExit(alcotestConfig)
	logger.Global(logger.New(loggerConfig))
	defer logger.Close()

	helloHandler := http.StripPrefix(helloPath, hello.Handler(helloConfig))
	dumpHandler := http.StripPrefix(dumpPath, dump.Handler())

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, helloPath) {
			helloHandler.ServeHTTP(w, r)
			return
		}
		if strings.HasPrefix(r.URL.Path, dumpPath) {
			dumpHandler.ServeHTTP(w, r)
			return
		}

		w.WriteHeader(http.StatusNotFound)
	})

	httputils.New(serverConfig).ListenAndServe(handler, nil, prometheus.New(prometheusConfig).Middleware, owasp.New(owaspConfig).Middleware, cors.New(corsConfig).Middleware)
}
