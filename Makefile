test:
	go test -v -cover ./...

run:
	go run main.go

proto-api:
	rm -rf internal/pb/*.go 2>/dev/null
	protoc --go_out=internal/pb --go_opt=paths=source_relative --go-grpc_out=internal/pb --go-grpc_opt=paths=source_relative --grpc-gateway_out=internal/pb --grpc-gateway_opt=paths=source_relative --proto_path=../bgmood-protos/api-gateway ../bgmood-protos/api-gateway/*.proto

proto-auth:
	rm -rf internal/authpb/* 2>/dev/null
	protoc --go_out=internal/authpb --go_opt=paths=source_relative --go-grpc_out=internal/authpb --go-grpc_opt=paths=source_relative --grpc-gateway_out=internal/authpb --grpc-gateway_opt=paths=source_relative --proto_path=../bgmood-protos/auth-service ../bgmood-protos/auth-service/*.proto

proto-circles:
	rm -rf internal/cpb/* 2>/dev/null
	protoc --go_out=internal/cpb --go_opt=paths=source_relative --go-grpc_out=internal/cpb --go-grpc_opt=paths=source_relative --grpc-gateway_out=internal/cpb --grpc-gateway_opt=paths=source_relative --proto_path=../bgmood-protos/circles-service ../bgmood-protos/circles-service/*.proto

proto:
	make proto-auth
	make proto-api
	make proto-circles

.PHONY: test run proto proto-auth proto-api proto-circles