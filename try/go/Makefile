.PHONY: test

test:
	go test -v ./...

codegen:
	protoc -I. --go_out=. proto/*.proto 
	protoc -I. --go_out=. --go-grpc_out=. proto/*.proto