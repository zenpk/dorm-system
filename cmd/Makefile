main:
	go build -o ../bin/main main.go
grpc-all:
	make -C dorm grpc
	make -C order grpc
	make -C team grpc
	make -C token grpc
	make -C user grpc
build-all:
	make -C dorm build
	make -C order build
	make -C team build
	make -C token build
	make -C user build
	make main
build-win-all:
	make -C dorm build-win
	make -C order build-win
	make -C team build-win
	make -C token build-win
	make -C user build-win
