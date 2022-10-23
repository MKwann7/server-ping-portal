FROM golang:1.18-stretch

ENV GO111MODULE=on

# install build essentials
RUN apt-get update && \
    apt-get install -y wget build-essential pkg-config --no-install-recommends

# Setup document root
RUN mkdir -p /app/code
RUN mkdir -p /app/logs
RUN mkdir -p /app/k6
RUN mkdir -p /app/tmp

WORKDIR /app/code

COPY app/core/go.mod /app/code
COPY app/core/go.sum /app/code
COPY app/core/reflex.conf /app/code

RUN go mod download

RUN ["go", "install", "github.com/cespare/reflex@latest"]

COPY app/core/src src
COPY docker/env/ping-local.env .env

EXPOSE 8080

ENTRYPOINT ["reflex", "-c", "./reflex.conf"]