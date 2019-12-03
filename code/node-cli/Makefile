.Phony: build
build:
	CGO_ENABLED=0 go build -o bin/node-cli ./cmd/
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/linux-node-cli ./cmd/


.Phony: compress
compress:
	CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/node-cli ./cmd/
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/linux-node-cli ./cmd/
	upx bin/node-cli
	upx bin/linux-node-cli