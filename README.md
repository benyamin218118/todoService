# Todo Service - ICE Global Golang Assignment

This is a Todo Service built in Go following **Clean Architecture** principles.  
It uses **MySQL**, **Redis Stream**, and **S3 (LocalStack)**, and provides REST endpoints to manage Todo items and file uploads.

---

## Features

- Manage `TodoItem` entities:
  - `id` (UUID)
  - `description` (string)
  - `dueDate` (timestamp)
  - `fileId` (string) → reference to file in S3
- Upload files to S3 via `/upload` endpoint
- Create Todo items via `/todo` endpoint
- Publish created Todo items to Redis Stream
- Fully containerized with Docker Compose
- Unit tests with mocks for MySQL, Redis, and S3
- Basic benchmarks for inserts, file upload, and Redis publishing
- Swagger API documentation

---

## Prerequisites

- Docker  
- Docker Compose  
- Make  
- Go >= 1.25.4  

---

## Setup

1. Copy `.env.example` to `.env` and update if necessary:

```env
LISTEN_HOST=0.0.0.0
LISTEN_PORT=8080
DB_DSN=dbuser:thepass@tcp(172.17.0.1:3306)/todo
REDIS_URL=redis://172.17.0.1:6379
S3_URL=http://172.17.0.1:4566
S3_BUCKET=todoservicefiles
S3_ACCESSKEY=TEST
S3_SECRETKEY=TEST
```

2. Build and run services:

```bash
make run
# or
make run-local
```

This will start:

- Go app on port `8080`
- MySQL database
- Redis
- LocalStack (S3)

3. Apply database migrations:

```bash
make migrate
# or
make migrate-local
```

---

## Endpoints

### 1. Upload File

```
POST /upload
Content-Type: multipart/form-data
Body: file=<file>
```

- Returns JSON:

```json
{
  "fileId": "generated-file-id"
}
```

- Validations:
  - Allowed file types: `.jpg`, `.png`, `.txt`
  - Max file size: 10MB

---

### 2. Create Todo Item

```
POST /todo
Content-Type: application/json
Body:
{
  "description": "Buy groceries",
  "dueDate": "2025-11-15T12:00:00Z",
  "fileId": "optional-file-id"
}
```

- Returns created Todo item:

```json
{
  "id": "uuid",
  "description": "Buy groceries",
  "dueDate": "2025-11-15T12:00:00Z",
  "fileId": "optional-file-id"
}
```

- Pushes the item to Redis stream `todo_stream`.

---

### 3. Swagger

Swagger docs are available in the `docs` folder:

- `docs/swagger.yaml`  
- `docs/swagger.json`  

You can view it using Swagger UI or any online swagger viewer or launch the service and browser http://localhost:8080/swagger/index.html

---

## Running Tests

Unit tests:

```bash
make test
# or
make test-local
```

This runs tests with **mocked MySQL, Redis, and S3**.

---

## Running Benchmarks

Benchmarks for core operations:

```bash
make bench
#or
make bench-local
```

Benchmarks include:

- Inserting a Todo item
- Uploading a file to S3
- Publishing messages to Redis stream

---

## Project Structure

```
.
├── app.go
├── docker-compose.yml
├── Dockerfile
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── domain
│   ├── config.go
│   ├── contracts
│   ├── error.go
│   └── todo.go
├── go.mod
├── go.sum
├── infra
│   ├── config
│   ├── db
│   ├── delivery
│   └── repositories
├── interface
│   └── controller
├── main.go
├── Makefile
├── README.md
├── todoService
└── usecase
```

---

## Makefile Commands

- `make build` → Build locally
- `make run` → Build and start all services using docker compose
- `make migrate` → Apply database migrations in docker
- `make test` → Run unit tests in a docker container
- `make benchmark` → Run benchmarks in a docker container

you can run them locally by adding a -local to the end of the command like:
`make migrate-local`

---

## Notes

- Followed Clean Architecture: **Domain → Use Cases → Interfaces → Infrastructure**.  
- Assignment is for technical evaluation only; not production-ready.
- This implementation strictly follows the assignment requirements and nothing extra was added.

