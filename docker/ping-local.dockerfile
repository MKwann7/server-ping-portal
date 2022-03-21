FROM golang:1.18-stretch

ENV GO111MODULE=on

# Setup document root
RUN mkdir -p /app/code
RUN mkdir -p /app/logs
RUN mkdir -p /app/k6
RUN mkdir -p /app/tmp

WORKDIR /app/code

COPY go.mod /app/code
COPY go.sum /app/code
COPY reflex.conf /app/code

RUN go mod download

RUN ["go", "install", "github.com/cespare/reflex@latest"]

COPY src src
COPY docker/env/ping-local.env .env
COPY code.iml code.iml

EXPOSE 8080

ENTRYPOINT ["reflex", "-c", "./reflex.conf"]