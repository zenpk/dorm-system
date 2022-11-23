build:
	go build -o ../../bin/token
build-win:
	go build -o ../../bin/token.exe
grpc:
	protoc token.proto --proto_path=../../internal/proto \
	--go_out=../../internal/service/token --go_opt=paths=source_relative \
	--go-grpc_out=../../internal/service/token --go-grpc_opt=paths=source_relative
