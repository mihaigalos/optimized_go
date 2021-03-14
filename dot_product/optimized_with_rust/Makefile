ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

build:
	cd lib && cargo build --release
	mv lib/target/release/libdotproduct.so lib/
	rm -rf lib/target
	go build -ldflags="-r $(ROOT_DIR)lib" main.go

run: build
	./main
