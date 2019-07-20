# goweb

[![Build Status](https://travis-ci.org/ViBiOh/goweb.svg?branch=master)](https://travis-ci.org/ViBiOh/goweb)
[![codecov](https://codecov.io/gh/ViBiOh/goweb/branch/master/graph/badge.svg)](https://codecov.io/gh/ViBiOh/goweb)
[![Go Report Card](https://goreportcard.com/badge/github.com/ViBiOh/goweb)](https://goreportcard.com/report/github.com/ViBiOh/goweb)

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
  -cert string
        [http] Certificate file
  -corsCredentials
        [cors] Access-Control-Allow-Credentials
  -corsExpose string
        [cors] Access-Control-Expose-Headers
  -corsHeaders string
        [cors] Access-Control-Allow-Headers (default "Content-Type")
  -corsMethods string
        [cors] Access-Control-Allow-Methods (default "GET")
  -corsOrigin string
        [cors] Access-Control-Allow-Origin (default "*")
  -csp string
        [owasp] Content-Security-Policy (default "default-src 'self'; base-uri 'self'")
  -frameOptions string
        [owasp] X-Frame-Options (default "deny")
  -graceful string
        [http] Graceful close duration (default "35s")
  -hsts
        [owasp] Indicate Strict Transport Security (default true)
  -key string
        [http] Key file
  -location string
        [hello] TimeZone for displaying current time (default "Europe/Paris")
  -port int
        [http] Listen port (default 1080)
  -prometheusPath string
        [prometheus] Path for exposing metrics (default "/metrics")
  -tracingAgent string
        [tracing] Jaeger Agent (e.g. host:port) (default "jaeger:6831")
  -tracingName string
        [tracing] Service name
  -url string
        [alcotest] URL to check
  -userAgent string
        [alcotest] User-Agent for check (default "Golang alcotest")
```
