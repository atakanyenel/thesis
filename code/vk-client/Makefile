GO111MODULE := on
export GO111MODULE

.PHONY: build
build: bin/virtual-kubelet

.PHONY: clean
clean: files := bin/virtual-kubelet bin/linux-virtual-kubelet
clean:
	@rm $(files) &>/dev/null || exit 0

.PHONY: mod
mod:
	@go mod tidy

bin/virtual-kubelet: BUILD_VERSION          ?= $(shell git describe --tags --always --dirty="-dev")
bin/virtual-kubelet: BUILD_DATE             ?= $(shell date -u '+%Y-%m-%d-%H:%M UTC')
bin/virtual-kubelet: VERSION_FLAGS    := -ldflags='-X "main.buildVersion=$(BUILD_VERSION)" -X "main.buildTime=$(BUILD_DATE)"'

bin/%:
	CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' -o bin/$(*) $(VERSION_FLAGS) ./cmd/$(*)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-extldflags "-static"' -o bin/linux-$(*) $(VERSION_FLAGS) ./cmd/$(*)

run:
	bin/virtual-kubelet --nodename vk

dist:
	make build
	scp bin/linux-virtual-kubelet atakanyenel@stuserver:.

.PHONY: docker
docker:
	docker build -t atakanyenel/vk-client .