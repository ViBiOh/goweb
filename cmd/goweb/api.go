package main

import (
	"embed"
	"flag"
	"net/http"
	"os"
	"strings"

	"github.com/ViBiOh/goweb/pkg/delay"
	"github.com/ViBiOh/goweb/pkg/dump"
	"github.com/ViBiOh/goweb/pkg/hello"
	"github.com/ViBiOh/httputils/v4/pkg/alcotest"
	"github.com/ViBiOh/httputils/v4/pkg/cors"
	"github.com/ViBiOh/httputils/v4/pkg/flags"
	"github.com/ViBiOh/httputils/v4/pkg/health"
	"github.com/ViBiOh/httputils/v4/pkg/httputils"
	"github.com/ViBiOh/httputils/v4/pkg/logger"
	"github.com/ViBiOh/httputils/v4/pkg/owasp"
	"github.com/ViBiOh/httputils/v4/pkg/prometheus"
	"github.com/ViBiOh/httputils/v4/pkg/renderer"
	"github.com/ViBiOh/httputils/v4/pkg/server"
)

const (
	helloPath = "/hello"
	dumpPath  = "/dump"
	delayPath = "/delay"
)

//go:embed templates static
var content embed.FS

func main() {
	fs := flag.NewFlagSet("api", flag.ExitOnError)

	appServerConfig := server.Flags(fs, "")
	promServerConfig := server.Flags(fs, "prometheus", flags.NewOverride("Port", 9090), flags.NewOverride("IdleTimeout", "10s"), flags.NewOverride("ShutdownTimeout", "5s"))
	healthConfig := health.Flags(fs, "")

	alcotestConfig := alcotest.Flags(fs, "")
	loggerConfig := logger.Flags(fs, "logger")
	prometheusConfig := prometheus.Flags(fs, "prometheus")
	owaspConfig := owasp.Flags(fs, "")
	corsConfig := cors.Flags(fs, "cors")
	rendererConfig := renderer.Flags(fs, "", flags.NewOverride("PublicURL", "https://api.vibioh.fr"), flags.NewOverride("Title", "I'm a teapot 🫖"))

	helloConfig := hello.Flags(fs, "")

	logger.Fatal(fs.Parse(os.Args[1:]))

	alcotest.DoAndExit(alcotestConfig)
	logger.Global(logger.New(loggerConfig))
	defer logger.Close()

	appServer := server.New(appServerConfig)
	promServer := server.New(promServerConfig)
	prometheusApp := prometheus.New(prometheusConfig)
	healthApp := health.New(healthConfig)

	rendererApp, err := renderer.New(rendererConfig, content, nil)
	logger.Fatal(err)

	helloHandler := http.StripPrefix(helloPath, hello.Handler(helloConfig))
	dumpHandler := http.StripPrefix(dumpPath, dump.Handler())
	delayHandler := http.StripPrefix(delayPath, delay.Handler())
	rendererHandler := rendererApp.Handler(func(r *http.Request) (string, int, map[string]interface{}, error) {
		return "public", http.StatusTeapot, nil, nil
	})

	appHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, helloPath) {
			helloHandler.ServeHTTP(w, r)
			return
		}
		if strings.HasPrefix(r.URL.Path, dumpPath) {
			dumpHandler.ServeHTTP(w, r)
			return
		}
		if strings.HasPrefix(r.URL.Path, delayPath) {
			delayHandler.ServeHTTP(w, r)
			return
		}

		rendererHandler.ServeHTTP(w, r)
	})

	go promServer.Start("prometheus", healthApp.End(), prometheusApp.Handler())
	go appServer.Start("http", healthApp.End(), httputils.Handler(appHandler, healthApp, prometheusApp.Middleware, owasp.New(owaspConfig).Middleware, cors.New(corsConfig).Middleware))

	healthApp.WaitForTermination(appServer.Done())
	server.GracefulWait(appServer.Done(), promServer.Done())
}
