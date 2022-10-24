build:
	go build -o ../../bin/user main.go
build-win:
	go build -o ../../bin/user.exe main.go
grpc:
	protoc user.proto --proto_path=../../internal/idl \
	--go_out=../../internal/service/user --go_opt=paths=source_relative \
	--go-grpc_out=../../internal/service/user --go-grpc_opt=paths=source_relative
