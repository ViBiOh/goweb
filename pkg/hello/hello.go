package hello

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/ViBiOh/flags"
	"github.com/ViBiOh/httputils/v4/pkg/httpjson"
	"github.com/ViBiOh/httputils/v4/pkg/telemetry"
)

type Hello struct {
	Name string `json:"greeting"`
}

type Config struct {
	LocationName string
}

func Flags(fs *flag.FlagSet, prefix string) *Config {
	var config Config

	flags.New("Location", "TimeZone for displaying current time").Prefix(prefix).DocPrefix("hello").StringVar(fs, &config.LocationName, "Europe/Paris", nil)

	return &config
}

func Handler(config *Config) http.Handler {
	location, err := time.LoadLocation(config.LocationName)
	if err != nil {
		slog.LogAttrs(context.Background(), slog.LevelError, "loading location", slog.String("name", config.LocationName), slog.Any("error", err))
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		telemetry.SetRouteTag(ctx, "/hello")

		name := r.PathValue("name")
		if len(name) == 0 {
			name = "World"
		}

		httpjson.Write(ctx, w, http.StatusOK, Hello{fmt.Sprintf("Hello %s, current time in %s is %v !", name, location.String(), time.Now().In(location))})
	})
}
