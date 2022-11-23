build:
	go build -o ../../bin/user
build-win:
	go build -o ../../bin/user.exe
grpc:
	protoc user.proto --proto_path=../../internal/proto \
	--go_out=../../internal/service/user --go_opt=paths=source_relative \
	--go-grpc_out=../../internal/service/user --go-grpc_opt=paths=source_relative
