package main

import (
	"flag"
	"os"
	"time"

	"github.com/ViBiOh/flags"
	"github.com/ViBiOh/goweb/pkg/hello"
	"github.com/ViBiOh/httputils/v4/pkg/alcotest"
	"github.com/ViBiOh/httputils/v4/pkg/cors"
	"github.com/ViBiOh/httputils/v4/pkg/health"
	"github.com/ViBiOh/httputils/v4/pkg/logger"
	"github.com/ViBiOh/httputils/v4/pkg/owasp"
	"github.com/ViBiOh/httputils/v4/pkg/prometheus"
	"github.com/ViBiOh/httputils/v4/pkg/server"
	"github.com/ViBiOh/httputils/v4/pkg/tracer"
)

type configuration struct {
	appServer  server.Config
	promServer server.Config
	health     health.Config
	alcotest   alcotest.Config
	logger     logger.Config
	tracer     tracer.Config
	prometheus prometheus.Config
	owasp      owasp.Config
	cors       cors.Config
	hello      hello.Config
}

func newConfig() (configuration, error) {
	fs := flag.NewFlagSet("api", flag.ExitOnError)

	return configuration{
		appServer:  server.Flags(fs, ""),
		promServer: server.Flags(fs, "prometheus", flags.NewOverride("Port", uint(9090)), flags.NewOverride("IdleTimeout", 10*time.Second), flags.NewOverride("ShutdownTimeout", 5*time.Second)),
		health:     health.Flags(fs, ""),
		alcotest:   alcotest.Flags(fs, ""),
		logger:     logger.Flags(fs, "logger"),
		tracer:     tracer.Flags(fs, "tracer"),
		prometheus: prometheus.Flags(fs, "prometheus", flags.NewOverride("Gzip", false)),
		owasp:      owasp.Flags(fs, ""),
		cors:       cors.Flags(fs, "cors"),

		hello: hello.Flags(fs, ""),
	}, fs.Parse(os.Args[1:])
}
