package main

import (
	"context"
	"fmt"

	"github.com/ViBiOh/httputils/v4/pkg/health"
	"github.com/ViBiOh/httputils/v4/pkg/logger"
	"github.com/ViBiOh/httputils/v4/pkg/request"
	"github.com/ViBiOh/httputils/v4/pkg/telemetry"
)

type client struct {
	telemetry telemetry.Service
	health    *health.Service
}

func newClient(ctx context.Context, config configuration) (client, error) {
	var output client
	var err error

	logger.Init(config.logger)

	output.telemetry, err = telemetry.New(ctx, config.telemetry)
	if err != nil {
		return output, fmt.Errorf("telemetry: %w", err)
	}

	request.AddOpenTelemetryToDefaultClient(output.telemetry.MeterProvider(), output.telemetry.TracerProvider())

	output.health = health.New(config.health)

	return output, nil
}

func (c client) Close(ctx context.Context) {
	c.telemetry.Close(ctx)
}
