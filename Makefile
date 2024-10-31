proto:
	protoc \
	  --go_out=./pkg/pb \
	  --go_opt=paths=source_relative \
	  --go-grpc_out=./pkg/pb \
	  --go-grpc_opt=paths=source_relative \
	  --proto_path=./pkg/proto \
	  ./**/proto/*.proto

.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

build: vet
	go build person.go
.PHONY:build