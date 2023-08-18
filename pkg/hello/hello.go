package hello

import (
	"flag"
	"fmt"
	"html"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/ViBiOh/flags"
	"github.com/ViBiOh/httputils/v4/pkg/httpjson"
)

type Hello struct {
	Name string `json:"greeting"`
}

type Config struct {
	locationName *string
}

func Flags(fs *flag.FlagSet, prefix string) Config {
	return Config{
		locationName: flags.New("Location", "TimeZone for displaying current time").Prefix(prefix).DocPrefix("hello").String(fs, "Europe/Paris", nil),
	}
}

func Handler(config Config) http.Handler {
	location, err := time.LoadLocation(*config.locationName)
	if err != nil {
		slog.Error("loading location", "err", err, "name", *config.locationName)
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		name := strings.TrimPrefix(html.EscapeString(r.URL.Path), "/")
		if name == "" {
			name = "World"
		}

		httpjson.Write(w, http.StatusOK, Hello{fmt.Sprintf("Hello %s, current time in %s is %v !", name, location.String(), time.Now().In(location))})
	})
}
