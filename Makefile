.PHONY: build

default: build;

build:
	protoc -I protobufs/ protobufs/*.proto --go_out=plugins=grpc:metrics_service


