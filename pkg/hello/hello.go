package hello

import (
	"flag"
	"fmt"
	"html"
	"net/http"
	"strings"
	"time"

	"github.com/ViBiOh/httputils/pkg/httperror"
	"github.com/ViBiOh/httputils/pkg/httpjson"
	"github.com/ViBiOh/httputils/pkg/logger"
	"github.com/ViBiOh/httputils/pkg/tools"
)

// Hello represents the outputted welcome message
type Hello struct {
	Name string `json:"greeting"`
}

// Config of package
type Config struct {
	locationName *string
}

// Flags adds flags for configuring package
func Flags(fs *flag.FlagSet, prefix string) Config {
	return Config{
		locationName: fs.String(tools.ToCamel(fmt.Sprintf("%sLocation", prefix)), "Europe/Paris", "[hello] TimeZone for displaying current time"),
	}
}

// Handler for Hello request. Should be use with net/http
func Handler(config Config) http.Handler {
	location, err := time.LoadLocation(*config.locationName)
	if err != nil {
		logger.Error("error while loading location %s: %v", *config.locationName, err)
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

		if err := httpjson.ResponseJSON(w, http.StatusOK, Hello{fmt.Sprintf("Hello %s, current time in %s is %v !", name, location.String(), time.Now().In(location))}, httpjson.IsPretty(r)); err != nil {
			httperror.InternalServerError(w, err)
		}
	})
}
