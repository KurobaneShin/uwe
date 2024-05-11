run: build
	@./bin/uwe

build:
	@go build -o bin/uwe .

test:
	@go test ./... -v

db-status:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING="postgres://postgres:foobarbaz@localhost:5432?sslmode=disable" goose -dir=db/migrations status

db-up:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING="postgres://postgres:foobarbaz@localhost:5432?sslmode=disable" goose -dir=db/migrations up

# docker-up:
#  @docker run --name postgres -e POSTGRES_PASSWORD=foobarbaz -p 5432:5432 -d postgres 
