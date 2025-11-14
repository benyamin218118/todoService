FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o todo-service

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/todo-service .
RUN apk add --no-cache ca-certificates

CMD ["./todo-service"]
