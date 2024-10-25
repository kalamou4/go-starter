.PHONY: run build test migrate-up migrate-down

run:
	go run cmd/api/main.go

build:
	go build -o bin/api cmd/api/main.go

test:
	go test ./...

migrate-up:
	migrate -path internal/database/migrations -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:5432/${DB_NAME}?sslmode=disable" up

migrate-down:
	migrate -path internal/database/migrations -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:5432/${DB_NAME}?sslmode=disable" down
