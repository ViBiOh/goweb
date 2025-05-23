# goweb

[![Build](https://github.com/ViBiOh/goweb/workflows/Build/badge.svg)](https://github.com/ViBiOh/goweb/actions)

## Getting started

Golang binary is built with static link. You can download it directly from the [GitHub Release page](https://github.com/ViBiOh/goweb/releases) or build it by yourself by cloning this repo and running `make`.

You can configure app by passing CLI args or environment variables (cf. [Usage](#usage) section). CLI override environment variables.

You'll find a Kubernetes exemple in the [`infra/`](infra) folder, using my [`app chart`](https://github.com/ViBiOh/charts/tree/main/app).

## Usage

The application can be configured by passing CLI args described below or their equivalent as environment variable. CLI values take precedence over environments variables.

Be careful when using the CLI values, if someone list the processes on the system, they will appear in plain-text. Pass secrets by environment variables: it's less easily visible.

```bash
Usage of api:
  --address           string    [server] Listen address ${API_ADDRESS}
  --cert              string    [server] Certificate file ${API_CERT}
  --corsCredentials             [cors] Access-Control-Allow-Credentials ${API_CORS_CREDENTIALS} (default false)
  --corsExpose        string    [cors] Access-Control-Expose-Headers ${API_CORS_EXPOSE}
  --corsHeaders       string    [cors] Access-Control-Allow-Headers ${API_CORS_HEADERS} (default "Content-Type")
  --corsMethods       string    [cors] Access-Control-Allow-Methods ${API_CORS_METHODS} (default "GET")
  --corsOrigin        string    [cors] Access-Control-Allow-Origin ${API_CORS_ORIGIN} (default "*")
  --csp               string    [owasp] Content-Security-Policy ${API_CSP} (default "default-src 'self'; base-uri 'self'")
  --frameOptions      string    [owasp] X-Frame-Options ${API_FRAME_OPTIONS} (default "deny")
  --graceDuration     duration  [http] Grace duration when signal received ${API_GRACE_DURATION} (default 30s)
  --hsts                        [owasp] Indicate Strict Transport Security ${API_HSTS} (default true)
  --idleTimeout       duration  [server] Idle Timeout ${API_IDLE_TIMEOUT} (default 2m0s)
  --key               string    [server] Key file ${API_KEY}
  --location          string    [hello] TimeZone for displaying current time ${API_LOCATION} (default "Europe/Paris")
  --loggerJson                  [logger] Log format as JSON ${API_LOGGER_JSON} (default false)
  --loggerLevel       string    [logger] Logger level ${API_LOGGER_LEVEL} (default "INFO")
  --loggerLevelKey    string    [logger] Key for level in JSON ${API_LOGGER_LEVEL_KEY} (default "level")
  --loggerMessageKey  string    [logger] Key for message in JSON ${API_LOGGER_MESSAGE_KEY} (default "msg")
  --loggerTimeKey     string    [logger] Key for timestamp in JSON ${API_LOGGER_TIME_KEY} (default "time")
  --name              string    [server] Name ${API_NAME} (default "http")
  --okStatus          int       [http] Healthy HTTP Status code ${API_OK_STATUS} (default 204)
  --port              uint      [server] Listen port (0 to disable) ${API_PORT} (default 1080)
  --pprofAgent        string    [pprof] URL of the Datadog Trace Agent (e.g. http://datadog.observability:8126) ${API_PPROF_AGENT}
  --pprofPort         int       [pprof] Port of the HTTP server (0 to disable) ${API_PPROF_PORT} (default 0)
  --readTimeout       duration  [server] Read Timeout ${API_READ_TIMEOUT} (default 5s)
  --shutdownTimeout   duration  [server] Shutdown Timeout ${API_SHUTDOWN_TIMEOUT} (default 10s)
  --telemetryRate     string    [telemetry] OpenTelemetry sample rate, 'always', 'never' or a float value ${API_TELEMETRY_RATE} (default "always")
  --telemetryURL      string    [telemetry] OpenTelemetry gRPC endpoint (e.g. otel-exporter:4317) ${API_TELEMETRY_URL}
  --telemetryUint64             [telemetry] Change OpenTelemetry Trace ID format to an unsigned int 64 ${API_TELEMETRY_UINT64} (default true)
  --url               string    [alcotest] URL to check ${API_URL}
  --userAgent         string    [alcotest] User-Agent for check ${API_USER_AGENT} (default "Alcotest")
  --writeTimeout      duration  [server] Write Timeout ${API_WRITE_TIMEOUT} (default 10s)
```
