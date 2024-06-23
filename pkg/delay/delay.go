package delay

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ViBiOh/httputils/v4/pkg/telemetry"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	telemetry.SetRouteTag(r.Context(), "/delay")

	path := strings.Trim(r.URL.Path, "/")
	delay, err := strconv.ParseInt(path, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	time.Sleep(time.Duration(delay) * time.Second)
}
