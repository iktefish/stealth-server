build:
	@go build -o bin/stealth-server

run: build
	@./bin/stealth-server

test:
	@go test -v ./...
