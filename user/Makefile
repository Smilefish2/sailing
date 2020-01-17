
GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto:
# --gofast_out=./gofast --gogo_out=./gogo
	protoc --proto_path=${GOPATH}/src:. --go_out=. --micro_out=. --gorm_out=.  proto/user/user.proto

.PHONY: build
build: proto

	go build -o user-srv main.go plugin.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t user-srv:latest
