run: build
	@./bin/uwe

build:
	@go build -o bin/uwe .
test:
	@go test ./... -v
