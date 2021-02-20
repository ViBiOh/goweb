package dump

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"github.com/ViBiOh/httputils/v4/pkg/httperror"
	"github.com/ViBiOh/httputils/v4/pkg/logger"
	"github.com/ViBiOh/httputils/v4/pkg/request"
)

// Handler for dump request. Should be use with net/http
func Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		value := dumpRequest(r)

		logger.Info("Dump of request\n%s", value)

		if _, err := w.Write([]byte(value)); err != nil {
			httperror.InternalServerError(w, err)
		}
	})
}

func getBufferContent(content map[string][]string) bytes.Buffer {
	var buffer bytes.Buffer

	for key, values := range content {
		buffer.WriteString(fmt.Sprintf("%s: %s\n", key, strings.Join(values, ",")))
	}

	return buffer
}

func dumpRequest(r *http.Request) string {
	parts := map[string]bytes.Buffer{
		"Headers": getBufferContent(r.Header),
		"Params":  getBufferContent(r.URL.Query()),
	}

	var form bytes.Buffer
	if err := r.ParseForm(); err != nil {
		form.WriteString(err.Error())
	} else {
		parts["Form"] = getBufferContent(r.PostForm)
	}

	body, err := request.ReadBodyRequest(r)
	if err != nil {
		logger.Error("%s", err)
	}

	var outputPattern bytes.Buffer
	outputPattern.WriteString("%s %s\n")
	outputData := []interface{}{
		r.Method,
		r.URL.Path,
	}

	for key, value := range parts {
		if value.Len() == 0 {
			continue
		}

		outputPattern.WriteString(key)
		outputPattern.WriteString("\n%s\n")
		outputData = append(outputData, value.String())
	}

	if len(body) != 0 {
		outputPattern.WriteString("Body\n%s\n")
		outputData = append(outputData, body)
	}

	return fmt.Sprintf(outputPattern.String(), outputData...)
}
