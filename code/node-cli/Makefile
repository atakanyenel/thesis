.Phony: build
build:
	CGO_ENABLED=0 go build -o bin/node-cli ./cmd/
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/linux-node-cli ./cmd/