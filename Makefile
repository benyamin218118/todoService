APP_NAME := app

.PHONY: all build run down migrate test bench

all: run

build:
	go get
	swag init -g main.go
	go build -o todoService .
run:
	swag init -g main.go
	docker compose up --build -d
	sleep 8s
	docker compose exec $(APP_NAME) ./todoService -run-migrations
test:
	docker compose exec $(APP_NAME) go test ./...

bench:
	docker compose exec $(APP_NAME) go test -bench=. ./...
down:
	docker compose down
migrate:
	docker compose exec $(APP_NAME) ./todoService -run-migrations

run-local:
	./todoService --run-migrations
	./todoService
	
migrate-local:
	./todoService -run-migrations
test-local:
	go test ./...
bench-local:
	go test -bench=. ./...