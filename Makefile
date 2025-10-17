MINIAUDIO_VERSION=$(shell git -C pkg/miniaudio describe --tags)

build-linux-amd64:
	mkdir -p tmp build
	zig cc -target x86_64-linux-gnu -c -fPIC pkg/main_linux.c -o tmp/libminiaudio-$(MINIAUDIO_VERSION)-linux-amd64.o -lpthread -lm -ldl
	zig cc -target x86_64-linux-gnu -shared tmp/libminiaudio-$(MINIAUDIO_VERSION)-linux-amd64.o -o build/libminiaudio-$(MINIAUDIO_VERSION)-linux-amd64.so

build-linux-arm64:
	mkdir -p tmp build
	zig cc -target aarch64-linux-gnu -c -fPIC pkg/main_linux.c -o tmp/libminiaudio-$(MINIAUDIO_VERSION)-linux-arm64.o -lpthread -lm -ldl
	zig cc -target aarch64-linux-gnu -shared tmp/libminiaudio-$(MINIAUDIO_VERSION)-linux-arm64.o -o build/libminiaudio-$(MINIAUDIO_VERSION)-linux-arm64.so

build-windows-amd64:
	mkdir -p tmp build
	zig cc -target x86_64-windows-gnu -c -fPIC pkg/main_windows.c -o tmp/libminiaudio-$(MINIAUDIO_VERSION)-windows-amd64.o
	zig cc -target x86_64-windows-gnu -shared tmp/libminiaudio-$(MINIAUDIO_VERSION)-windows-amd64.o -o build/libminiaudio-$(MINIAUDIO_VERSION)-windows-amd64.dll

build-windows-arm64:
	mkdir -p tmp build
	zig cc -target aarch64-windows-gnu -c -fPIC pkg/main_windows.c -o tmp/libminiaudio-$(MINIAUDIO_VERSION)-windows-arm64.o
	zig cc -target aarch64-windows-gnu -shared tmp/libminiaudio-$(MINIAUDIO_VERSION)-windows-arm64.o -o build/libminiaudio-$(MINIAUDIO_VERSION)-windows-arm64.dll

build-darwin-amd64:
	mkdir -p tmp build
	zig cc -target x86_64-macos-gnu -c -fPIC pkg/main_darwin.c -o tmp/libminiaudio-$(MINIAUDIO_VERSION)-darwin-amd64.o
	zig cc -target x86_64-macos-gnu -shared tmp/libminiaudio-$(MINIAUDIO_VERSION)-darwin-amd64.o -o build/libminiaudio-$(MINIAUDIO_VERSION)-darwin-amd64.dylib

build-darwin-arm64:
	mkdir -p tmp build
	zig cc -target aarch64-macos-gnu -c -fPIC pkg/main_darwin.c -o tmp/libminiaudio-$(MINIAUDIO_VERSION)-darwin-arm64.o
	zig cc -target aarch64-macos-gnu -shared tmp/libminiaudio-$(MINIAUDIO_VERSION)-darwin-arm64.o -o build/libminiaudio-$(MINIAUDIO_VERSION)-darwin-arm64.dylib

build-all: build-linux-amd64 build-linux-arm64 build-windows-amd64 build-windows-arm64

build:
	CGO_ENABLED=0 go build -o bin/ cmd/main.go

run:
	CGO_ENABLED=0 go run cmd/main.go
