.PHONY: build test clean docker prepare update

GO=CGO_ENABLED=0 GO111MODULE=on go

MICROSERVICES=example/cmd/device-simple/device-simple
.PHONY: $(MICROSERVICES)

VERSION=$(shell cat ./VERSION)

GOFLAGS=-ldflags "-X github.com/edgexfoundry/device-sdk-go.Version=$(VERSION)"

GIT_SHA=$(shell git rev-parse HEAD)

build: $(MICROSERVICES)
	$(GO) build ./...

example/cmd/device-simple/device-simple:
	$(GO) build $(GOFLAGS) -o $@ ./example/cmd/device-simple

docker:
	docker build \
		-f example/cmd/device-simple/Dockerfile \
		--label "git_sha=$(GIT_SHA)" \
		-t edgexfoundry/docker-device-sdk-simple:$(GIT_SHA) \
		-t edgexfoundry/docker-device-sdk-simple:$(VERSION)-dev \
		--build-arg http_proxy=${HTTP_PROXY} \
		--build-arg https_proxy=${HTTPS_PROXY} \
		.

test:
	$(GO) test ./... -cover

clean:
	rm -f $(MICROSERVICES)

prepare:
	glide install

update:
	glide update
