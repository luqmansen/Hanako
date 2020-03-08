
GOPATH:=$(shell go env GOPATH)

du: #short for docker up
	docker-compose -f docker-helper.yml up -d

dd: #short for docker down
	docker-compose -f docker-helper.yml down

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
