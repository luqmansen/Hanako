
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/anime/anime.proto
	sed -i "s|v2/||g" proto/anime/*.pb.micro.go


.PHONY: build
build: proto

	go build -o anime-srv *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t anime-srv:latest
