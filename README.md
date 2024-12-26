# portto-go-interview

## Prerequisites

### Docker

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

### Go (1.22.5^)

- [Go](https://golang.org/dl/)

## Configuration

1. copy .env.example to .env and fill in the values
2. copy docker-compose-example.yml to docker-compose.yml

## Run the server on docker

```bash
docker compose up -d --build
```

## Run the server on local

```bash
# run "docker compose down" if you want to clean up the container
docker compose -f docker-compose-local.yml up -d --build
cd Go-Service
go mod download
cd src/main/
go run main.go
```

## Test

### Integration test

```bash
cd Go-Service
go mod download
cd src/test/integration/
go test -v meme_coin_usecase_test.go
```

### End-to-end test

```bash
# run "docker compose down" if you want to clean up the container
docker compose -f docker-compose-local.yml up -d --build
cd Go-Service
go mod download
cd src/test/end_to_end/
go test -v meme_coin_test.go
```
