package main

import (
	"context"
	"fmt"

	"github.com/ViBiOh/httputils/v4/pkg/health"
	"github.com/ViBiOh/httputils/v4/pkg/logger"
	"github.com/ViBiOh/httputils/v4/pkg/prometheus"
	"github.com/ViBiOh/httputils/v4/pkg/request"
	"github.com/ViBiOh/httputils/v4/pkg/tracer"
)

type client struct {
	tracer     tracer.App
	logger     logger.Logger
	prometheus prometheus.App
	health     health.App
}

func newClient(ctx context.Context, config configuration) (client, error) {
	var output client
	var err error

	output.logger = logger.New(config.logger)
	logger.Global(output.logger)

	output.tracer, err = tracer.New(ctx, config.tracer)
	if err != nil {
		return output, fmt.Errorf("tracer: %w", err)
	}

	request.AddTracerToDefaultClient(output.tracer.GetProvider())

	output.prometheus = prometheus.New(config.prometheus)
	output.health = health.New(config.health)

	return output, nil
}

func (c client) Close(ctx context.Context) {
	c.tracer.Close(ctx)
	c.logger.Close()
}
