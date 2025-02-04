# Stage 1: build
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /bin/pinger main.go

FROM alpine:3.17

WORKDIR /app

COPY --from=builder /bin/pinger /app/pinger

RUN apk update && apk add --no-cache iputils

ENTRYPOINT ["/app/pinger"]
