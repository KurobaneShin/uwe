dsn = "postgres://postgres:foobarbaz@localhost:5432?sslmode=disable" 
migrationPath = "db/migrations" 

run: build
	@./bin/uwe

build:
	@go build -o bin/uwe .

test:
	@go test ./... -v

seed:
	@go run cmd/seed/main.go

db-status:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(dsn) goose -dir=$(migrationPath) status

db-up:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(dsn) goose -dir=$(migrationPath) up

db-reset:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(dsn) goose -dir=$(migrationPath) reset

migration-new:
	@GOOSE_MIGRATION_DIR=$(migrationPath) goose create $(filter-out $@,$(MAKECMDGOALS))

docker-up:
 docker run --name postgres -e POSTGRES_PASSWORD=foobarbaz -p 5432:5432 -d postgres 
