FROM golang:1.12 as builder

WORKDIR /app
COPY . .

RUN make

ARG CODECOV_TOKEN
RUN curl -s https://codecov.io/bash | bash

FROM alpine as fetcher

WORKDIR /app

RUN apk --update add curl \
 && curl -s -o /app/cacert.pem https://curl.haxx.se/ca/cacert.pem \
 && curl -s -o /app/zoneinfo.zip https://raw.githubusercontent.com/golang/go/master/lib/time/zoneinfo.zip

FROM scratch

ENV ZONEINFO zoneinfo.zip
EXPOSE 1080

HEALTHCHECK --retries=5 CMD [ "/goweb", "-url", "http://localhost:1080/health" ]
ENTRYPOINT [ "/goweb" ]

ARG APP_VERSION
ENV VERSION=${APP_VERSION}

COPY doc /doc
COPY --from=fetcher /app/cacert.pem /etc/ssl/certs/ca-certificates.crt
COPY --from=fetcher /app/zoneinfo.zip /
COPY --from=builder /app/bin/goweb /
