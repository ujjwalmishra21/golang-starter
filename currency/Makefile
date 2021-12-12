.PHONY: protos

protos:
	protoc --go_out=protos/currency --go_opt=paths=source_relative --go-grpc_out=protos/currency --go-grpc_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false protos/currency.proto -I protos/