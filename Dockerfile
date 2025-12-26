FROM rg.fr-par.scw.cloud/vibioh/scratch

ENV API_PORT 1080
EXPOSE 1080

ENV ZONEINFO /zoneinfo.zip
COPY zoneinfo.zip /zoneinfo.zip
COPY ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

HEALTHCHECK --retries=5 CMD [ "/goweb", "-url", "http://127.0.0.1:1080/health" ]
ENTRYPOINT [ "/goweb" ]

ARG VERSION
ENV VERSION ${VERSION}

ARG GIT_SHA
ENV GIT_SHA ${GIT_SHA}

ARG TARGETOS
ARG TARGETARCH

COPY release/goweb_${TARGETOS}_${TARGETARCH} /goweb
