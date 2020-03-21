
GOPATH:=$(shell go env GOPATH)

du: #short for docker up
	docker-compose -f docker-helper.yml up -d

.PHONY: dev
dev:
	go run main.go

.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/anime.proto

.PHONY: build
build: proto

	go build -o anime-srv *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t anime-srv:latest
