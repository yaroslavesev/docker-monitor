FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /bin/backend main.go

FROM alpine:3.17

WORKDIR /app

COPY --from=builder /bin/backend /app/backend

COPY .env /app/.env

EXPOSE 8080

ENTRYPOINT ["/app/backend"]
