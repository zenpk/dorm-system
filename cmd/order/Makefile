build:
	go build -o ../../bin/order
build-win:
	go build -o ../../bin/order.exe
grpc:
	protoc order.proto --proto_path=../../internal/api \
	--go_out=../../internal/service/order --go_opt=paths=source_relative \
	--go-grpc_out=../../internal/service/order --go-grpc_opt=paths=source_relative
