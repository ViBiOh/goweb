package dump

import (
	"bytes"
	"context"
	"fmt"
	"html"
	"io"
	"log/slog"
	"net/http"
	"sort"
	"strings"

	"github.com/ViBiOh/httputils/v4/pkg/httperror"
	"github.com/ViBiOh/httputils/v4/pkg/telemetry"
	"go.opentelemetry.io/otel/metric"
)

func Handler(meterProvider metric.MeterProvider) http.Handler {
	counter, err := meterProvider.Meter("github.com/ViBiOh/goweb").Int64Counter("goweb.dump")
	if err != nil {
		slog.LogAttrs(context.Background(), slog.LevelError, "create dump counter", slog.Any("error", err))
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		telemetry.SetRouteTag(ctx, "/dump")

		value, err := dumpRequest(r)
		if err != nil {
			httperror.BadRequest(ctx, w, err)
			return
		}

		counter.Add(ctx, 1)

		slog.LogAttrs(ctx, slog.LevelInfo, "Dump of request", slog.String("content", value))

		if _, err := w.Write([]byte(html.EscapeString(value))); err != nil {
			httperror.InternalServerError(ctx, w, err)
		}
	})
}

func dumpRequest(r *http.Request) (string, error) {
	parts := map[string]string{
		"Headers": getBufferContent(r.Header),
		"Params":  getBufferContent(r.URL.Query()),
		"Referer": r.Referer(),
	}

	if err := r.ParseForm(); err != nil {
		return "", fmt.Errorf("parse form: %w", err)
	}

	cookies := r.Cookies()
	cookiesString := make([]string, len(cookies))
	for i, cookie := range cookies {
		cookiesString[i] = cookie.String()
	}
	parts["Cookies"] = strings.Join(cookiesString, ", ")

	parts["Form"] = getBufferContent(r.PostForm)

	body, err := readContent(r.Body)
	if err != nil {
		return "", fmt.Errorf("read content: %w", err)
	}

	var outputPattern bytes.Buffer
	outputPattern.WriteString("RemoteAddr=`%s`\nHost=`%s`\n%s %s")
	outputData := []any{
		r.RemoteAddr,
		r.Host,
		r.Method,
		r.URL.Path,
	}

	for key, value := range parts {
		if len(value) == 0 {
			continue
		}

		outputPattern.WriteString("\n\n")
		outputPattern.WriteString(key)
		outputPattern.WriteString("\n%s")
		outputData = append(outputData, value)
	}

	if len(body) != 0 {
		outputPattern.WriteString("\nBody\n%s")
		outputData = append(outputData, body)
	}

	return fmt.Sprintf(outputPattern.String(), outputData...), nil
}

func getBufferContent(content map[string][]string) string {
	var output []string

	for key, values := range content {
		output = append(output, fmt.Sprintf("%s: %s", key, strings.Join(values, ",")))
	}

	sort.Strings(output)
	return strings.Join(output, "\n")
}

func readContent(body io.ReadCloser) (content []byte, err error) {
	if body == nil {
		return
	}

	defer func() {
		if closeErr := body.Close(); closeErr != nil {
			if err == nil {
				err = closeErr
			} else {
				err = fmt.Errorf("%s: %w", err, closeErr)
			}
		}
	}()

	content, err = io.ReadAll(body)
	return
}
