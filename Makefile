protoc:
	protoc -I . \
 	--go_out ./pkg/grpc/ \
 	--go_opt paths=source_relative  \
 	--go-grpc_out ./pkg/grpc/ \
 	--go-grpc_opt paths=source_relative \
 	-I./third_party/googleapis \
 	--grpc-gateway_out ./pkg/grpc \
 	--grpc-gateway_opt logtostderr=true \
 	--grpc-gateway_opt paths=source_relative \
 	--openapiv2_out ./api/ \
 	 mutual.proto

build-grpc:
	go build -o build ./cmd/mutual-grpc

build-http:
	go build -o build ./cmd/mutual-http

test-grpc:
	go run ./tests/test-client.go