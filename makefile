.PHONY: grpc-gen
grpc-gen:
	@echo "Generating gRPC files..."
	@protoc -I proto --go_out=./internal --go_opt=paths=source_relative --go-grpc_out=./internal  --go-grpc_opt=paths=source_relative proto/petstore/v1/*.proto

all: test build

.PHONY: test
test:
	@go test -v ./...

.PHONY: build
build: CGO_ENABLED=0
build:
	@go generate
	@go build -a -ldflags="-w -s -extldflags='-static'" -o bin/server main.go
	@echo "Binary is now available at bin/server"