.PHONY: proto

proto:
	for f in services/*/proto/*.proto; do \
		protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. $$f; \
        			echo compiled: $$f; \
	done