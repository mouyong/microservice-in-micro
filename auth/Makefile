
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/auth/auth.proto

.PHONY: build
build: proto

	go build -o auth-srv *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t auth-srv:latest
