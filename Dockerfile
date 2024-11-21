FROM golang:1.23-alpine AS build

ENV CGO_ENABLED=1

RUN apk add --no-cache gcc musl-dev

WORKDIR /app

ENTRYPOINT ./startup.sh
