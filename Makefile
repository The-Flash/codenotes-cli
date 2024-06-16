dev:
	go run ./cmd/cli/main.go
build:
	go build -o bin/codenotes cmd/cli/main.go