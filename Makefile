build:
	@go build -o bin/go-cray

test:
	@go test -v ./...

run: build
	@./bin/go-cray

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down