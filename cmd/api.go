package main

import (
	"flag"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/ViBiOh/goweb/pkg/hello"
	httputils "github.com/ViBiOh/httputils/v2/pkg"
	"github.com/ViBiOh/httputils/v2/pkg/alcotest"
	"github.com/ViBiOh/httputils/v2/pkg/cors"
	"github.com/ViBiOh/httputils/v2/pkg/logger"
	"github.com/ViBiOh/httputils/v2/pkg/opentracing"
	"github.com/ViBiOh/httputils/v2/pkg/owasp"
	"github.com/ViBiOh/httputils/v2/pkg/prometheus"
)

const (
	helloPath = "/hello"
	docPath   = "doc/"
)

func main() {
	fs := flag.NewFlagSet("api", flag.ExitOnError)

	serverConfig := httputils.Flags(fs, "")
	alcotestConfig := alcotest.Flags(fs, "")
	prometheusConfig := prometheus.Flags(fs, "prometheus")
	opentracingConfig := opentracing.Flags(fs, "tracing")
	owaspConfig := owasp.Flags(fs, "")
	corsConfig := cors.Flags(fs, "cors")

	helloConfig := hello.Flags(fs, "")

	logger.Fatal(fs.Parse(os.Args[1:]))

	alcotest.DoAndExit(alcotestConfig)

	prometheusApp := prometheus.New(prometheusConfig)
	opentracingApp := opentracing.New(opentracingConfig)
	owaspApp := owasp.New(owaspConfig)
	corsApp := cors.New(corsConfig)

	helloHandler := http.StripPrefix(helloPath, hello.Handler(helloConfig))

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, helloPath) {
			helloHandler.ServeHTTP(w, r)
			return
		}

		w.Header().Set("Cache-Control", "no-cache")
		http.ServeFile(w, r, path.Join(docPath, r.URL.Path))
	})

	httpHandler := httputils.ChainMiddlewares(handler, prometheusApp, opentracingApp, owaspApp, corsApp)

	httputils.New(serverConfig).ListenAndServe(httpHandler, httputils.HealthHandler(nil), nil)
}
