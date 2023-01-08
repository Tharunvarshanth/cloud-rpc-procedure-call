install-tool-grpc:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	go install github.com/micro/go-micro/cmd/protoc-gen-micro/v2

GOPATH=$(shell go env GOPATH)

generate-rpc-server-proto:
    protoc --proto_path=$(GOPATH)/src:. --micro_out=. --go_out=. serverRpcUser.proto

protoc --proto_path=/home/tharun/go:.  --proto_path=services/server/proto:. --micro_out=. --go_out=. serverRpcUser.proto

generate-rpc-server-protobuf:
	protoc --plugin=protoc-gen-go=$(GOPATH)/bin/protoc-gen-go --plugin=protoc-gen-micro=$(GOPATH)/bin/protoc-gen-micro --proto_path=services/server/proto:. --micro_out=. --go_out=. serverRpcUser.proto
