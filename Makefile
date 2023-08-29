test:
	go test -v -cover ./...

run:
	go run main.go

proto-gateway:
	rm -rf internal/pb/*.go 2>/dev/null
	protoc --go_out=internal/pb --go_opt=paths=source_relative --go-grpc_out=internal/pb --go-grpc_opt=paths=source_relative --grpc-gateway_out=internal/pb --grpc-gateway_opt=paths=source_relative --proto_path=internal/proto internal/proto/*.proto

proto-auth:
	rm -rf internal/authpb/* 2>/dev/null
	protoc --go_out=internal/authpb --go_opt=paths=source_relative --go-grpc_out=internal/authpb --go-grpc_opt=paths=source_relative --grpc-gateway_out=internal/authpb --grpc-gateway_opt=paths=source_relative --proto_path=../bgmood-protos/auth-service ../bgmood-protos/auth-service/*.proto

proto:
	make proto-auth
	make proto-gateway

.PHONY: test run proto proto-auth proto-gateway