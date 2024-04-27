package main

import (
	"flag"
	"os"

	"github.com/ViBiOh/flags"
	"github.com/ViBiOh/goweb/pkg/hello"
	"github.com/ViBiOh/httputils/v4/pkg/alcotest"
	"github.com/ViBiOh/httputils/v4/pkg/cors"
	"github.com/ViBiOh/httputils/v4/pkg/health"
	"github.com/ViBiOh/httputils/v4/pkg/logger"
	"github.com/ViBiOh/httputils/v4/pkg/owasp"
	"github.com/ViBiOh/httputils/v4/pkg/server"
	"github.com/ViBiOh/httputils/v4/pkg/telemetry"
)

type configuration struct {
	alcotest  *alcotest.Config
	telemetry *telemetry.Config
	hello     *hello.Config
	logger    *logger.Config
	cors      *cors.Config
	owasp     *owasp.Config
	appServer *server.Config
	health    *health.Config
}

func newConfig() configuration {
	fs := flag.NewFlagSet("api", flag.ExitOnError)
	fs.Usage = flags.Usage(fs)

	config := configuration{
		appServer: server.Flags(fs, ""),
		health:    health.Flags(fs, ""),
		alcotest:  alcotest.Flags(fs, ""),
		logger:    logger.Flags(fs, "logger"),
		telemetry: telemetry.Flags(fs, "telemetry"),
		owasp:     owasp.Flags(fs, ""),
		cors:      cors.Flags(fs, "cors"),

		hello: hello.Flags(fs, ""),
	}

	_ = fs.Parse(os.Args[1:])

	return config
}
