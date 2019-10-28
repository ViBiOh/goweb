# goweb

[![Build Status](https://travis-ci.org/ViBiOh/goweb.svg?branch=master)](https://travis-ci.org/ViBiOh/goweb)
[![codecov](https://codecov.io/gh/ViBiOh/goweb/branch/master/graph/badge.svg)](https://codecov.io/gh/ViBiOh/goweb)
[![Go Report Card](https://goreportcard.com/badge/github.com/ViBiOh/goweb)](https://goreportcard.com/report/github.com/ViBiOh/goweb)
[![Dependabot Status](https://api.dependabot.com/badges/status?host=github&repo=ViBiOh/goweb)](https://dependabot.com)

## CI

Following variables are required for CI:

| Name | Purpose |
|:--:|:--:|
| **GITHUB_OAUTH_TOKEN** | for creating release from  Github API |
| **DOMAIN** | for setting Traefik domain for app |
| **DEPLOY_CREDENTIALS** | for deploying app to server |
| **DOCKER_USER** | for publishing Docker image |
| **DOCKER_PASS** | for publishing Docker image |

## Usage

```bash
Usage of api:
  -address string
        [http] Listen address {API_ADDRESS}
  -cert string
        [http] Certificate file {API_CERT}
  -corsCredentials
        [cors] Access-Control-Allow-Credentials {API_CORS_CREDENTIALS}
  -corsExpose string
        [cors] Access-Control-Expose-Headers {API_CORS_EXPOSE}
  -corsHeaders string
        [cors] Access-Control-Allow-Headers {API_CORS_HEADERS} (default "Content-Type")
  -corsMethods string
        [cors] Access-Control-Allow-Methods {API_CORS_METHODS} (default "GET")
  -corsOrigin string
        [cors] Access-Control-Allow-Origin {API_CORS_ORIGIN} (default "*")
  -csp string
        [owasp] Content-Security-Policy {API_CSP} (default "default-src 'self'; base-uri 'self'")
  -frameOptions string
        [owasp] X-Frame-Options {API_FRAME_OPTIONS} (default "deny")
  -hsts
        [owasp] Indicate Strict Transport Security {API_HSTS} (default true)
  -key string
        [http] Key file {API_KEY}
  -location string
        [hello] TimeZone for displaying current time {API_LOCATION} (default "Europe/Paris")
  -port int
        [http] Listen port {API_PORT} (default 1080)
  -prometheusPath string
        [prometheus] Path for exposing metrics {API_PROMETHEUS_PATH} (default "/metrics")
  -url string
        [alcotest] URL to check {API_URL}
  -userAgent string
        [alcotest] User-Agent for check {API_USER_AGENT} (default "Golang alcotest")
```
