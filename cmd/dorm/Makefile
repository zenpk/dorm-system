build:
	go build -o ../../bin/dorm
build-win:
	go build -o ../../bin/dorm.exe
grpc:
	protoc dorm.proto --proto_path=../../internal/api \
	--go_out=../../internal/service/dorm --go_opt=paths=source_relative \
	--go-grpc_out=../../internal/service/dorm --go-grpc_opt=paths=source_relative
