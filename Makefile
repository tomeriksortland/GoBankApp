build:
	@go build -o bin/gobankapp

run: build
	@./bin/gobankapp

test:
	@go test -v ./...

