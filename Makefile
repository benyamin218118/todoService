APP_NAME := app

.PHONY: all build run down migrate test bench

all: run

build:
	go get
	swag init -g main.go
	go build -o todo-service .

run:
	swag init -g main.go
	docker compose up --build -d
	docker compose exec $(APP_NAME) ./todo-service -run-migrations


down:
	docker compose down

migrate:
	docker compose exec $(APP_NAME) ./todo-service -run-migrations

test:
	docker compose exec $(APP_NAME) go test ./...

bench:
	docker compose exec $(APP_NAME) go test -bench=. ./...