FROM golang:1.16 AS build

ADD . /src/

RUN cd /src/ && \
    CGO_ENABLED=0 go build -v -o /ts3exporter

FROM scratch

COPY --from=build --chown=100:100 /ts3exporter /ts3exporter

USER 1000:1000
EXPOSE 8080/tcp
ENTRYPOINT ["/ts3exporter"]