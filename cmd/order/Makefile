build:
	go build -o ../../bin/order main.go
build-win:
	go build -o ../../bin/order.exe main.go
grpc:
	protoc order.proto --proto_path=../../internal/idl \
	--go_out=../../internal/service/order --go_opt=paths=source_relative \
	--go-grpc_out=../../internal/service/order --go-grpc_opt=paths=source_relative
