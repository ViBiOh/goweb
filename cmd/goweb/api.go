package main

import (
	"context"
	"fmt"
	"net/http"

	_ "net/http/pprof"

	"github.com/ViBiOh/httputils/v4/pkg/alcotest"
	"github.com/ViBiOh/httputils/v4/pkg/cors"
	"github.com/ViBiOh/httputils/v4/pkg/httputils"
	"github.com/ViBiOh/httputils/v4/pkg/logger"
	"github.com/ViBiOh/httputils/v4/pkg/owasp"
	"github.com/ViBiOh/httputils/v4/pkg/recoverer"
	"github.com/ViBiOh/httputils/v4/pkg/server"
)

func main() {
	config, err := newConfig()
	if err != nil {
		logger.Fatal(fmt.Errorf("config: %s", err))
	}

	alcotest.DoAndExit(config.alcotest)

	go func() {
		fmt.Println(http.ListenAndServe("localhost:9999", http.DefaultServeMux))
	}()

	ctx := context.Background()

	client, err := newClient(ctx, config)
	if err != nil {
		logger.Fatal(fmt.Errorf("client: %s", err))
	}
	defer client.Close(ctx)

	appServer := server.New(config.appServer)
	promServer := server.New(config.promServer)

	go promServer.Start(client.health.End(ctx), "prometheus", client.prometheus.Handler())
	go appServer.Start(client.health.End(ctx), "http", httputils.Handler(newPort(config), client.health, recoverer.Middleware, client.prometheus.Middleware, client.tracer.Middleware, owasp.New(config.owasp).Middleware, cors.New(config.cors).Middleware))

	client.health.WaitForTermination(appServer.Done())
	server.GracefulWait(appServer.Done(), promServer.Done())
}
