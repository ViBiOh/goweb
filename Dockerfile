FROM alpine as fetcher

WORKDIR /app

RUN apk --update add curl \
 && curl -q -sSL --max-time 10 -o /app/cacert.pem https://curl.haxx.se/ca/cacert.pem \
 && curl -q -sSL --max-time 10 -o /app/zoneinfo.zip https://raw.githubusercontent.com/golang/go/master/lib/time/zoneinfo.zip

FROM scratch

ENV ZONEINFO zoneinfo.zip
EXPOSE 1080

HEALTHCHECK --retries=5 CMD [ "/goweb", "-url", "http://localhost:1080/health" ]
ENTRYPOINT [ "/goweb" ]

ARG APP_VERSION
ENV VERSION=${APP_VERSION}

ARG OS
ARG ARCH

COPY --from=fetcher /app/cacert.pem /etc/ssl/certs/ca-certificates.crt
COPY --from=fetcher /app/zoneinfo.zip /
COPY release/goweb_${OS}_${ARCH} /goweb
