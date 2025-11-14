# Builder
FROM golang:1.25.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o todoService

FROM debian:bookworm-slim
WORKDIR /app

RUN mkdir -p /app/infra/db
COPY --from=builder /app/infra/db/migrations /app/infra/db/migrations
COPY --from=builder /app/todoService .
COPY --from=builder /app/docs .

CMD ["./todoService"]