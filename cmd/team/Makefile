build:
	go build -o ../../bin/team
build-win:
	go build -o ../../bin/team.exe
grpc:
	protoc team.proto --proto_path=../../internal/idl \
	--go_out=../../internal/service/team --go_opt=paths=source_relative \
	--go-grpc_out=../../internal/service/team --go-grpc_opt=paths=source_relative
