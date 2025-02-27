build-linux-amd64:
	zig cc -target x86_64-linux-gnu -c -fPIC pkg/miniaudio/main_linux.c -o tmp/libminiaudio-linux-amd64.o -lpthread -lm -ldl
	zig cc -target x86_64-linux-gnu -shared tmp/libminiaudio-linux-amd64.o -o build/linux/libminiaudio-linux-amd64.so

build-linux-arm64:
	zig cc -target aarch64-linux-gnu -c -fPIC pkg/miniaudio/main_linux.c -o tmp/libminiaudio-linux-arm64.o -lpthread -lm -ldl
	zig cc -target aarch64-linux-gnu -shared tmp/libminiaudio-linux-arm64.o -o build/linux/libminiaudio-linux-arm64.so

build-windows-amd64:
	zig cc -target x86_64-windows-gnu -c -fPIC pkg/miniaudio/main_windows.c -o tmp/libminiaudio-windows-amd64.o
	zig cc -target x86_64-windows-gnu -shared tmp/libminiaudio-windows-amd64.o -o build/windows/libminiaudio-windows-amd64.dll

build-windows-arm64:
	zig cc -target aarch64-windows-gnu -c -fPIC pkg/miniaudio/main_windows.c -o tmp/libminiaudio-windows-arm64.o
	zig cc -target aarch64-windows-gnu -shared tmp/libminiaudio-windows-arm64.o -o build/windows/libminiaudio-windows-arm64.dll

build-all: build-linux-amd64 build-linux-arm64 build-windows-amd64 build-windows-arm64

build: CGO_ENABLED=0 go build -o bin/ cmd/main.go
