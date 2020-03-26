
GOPATH:=$(shell go env GOPATH)

.PHONY: api
api:
	MICRO_API_NAMESPACE=hanako.api micro api --handler=api

.PHONY: dev
dev:
	go run main.go

.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/rest.proto

.PHONY: build
build: proto

	go build -o rest-api *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t rest-api:latest
