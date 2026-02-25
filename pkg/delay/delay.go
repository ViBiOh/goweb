package delay

import (
	"net/http"
	"strconv"
	"time"

	"github.com/ViBiOh/httputils/v4/pkg/telemetry"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	telemetry.SetRouteTag(r.Context(), "/delay")

	duration := r.PathValue("duration")
	if len(duration) == 0 {
		duration = "1"
	}

	delay, err := strconv.ParseInt(duration, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	time.Sleep(time.Duration(delay) * time.Second)
}
