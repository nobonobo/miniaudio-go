//go:build linux || windows

package main

import (
	"fmt"
	"runtime"

	"github.com/ebitengine/purego"
)

func getSystemLibrary() string {
	switch runtime.GOOS {
	case "linux":
		switch runtime.GOARCH {
		case "amd64":
		case "arm64":
		default:
			panic(fmt.Errorf("GOARCH=%s is not supported", runtime.GOARCH))
		}
		return "libc.so.6"
	case "windows":
		return "ucrtbase.dll"
	default:
		panic(fmt.Errorf("GOOS=%s is not supported", runtime.GOOS))
	}
}

func main() {
	libc, err := openLibrary(getSystemLibrary())
	if err != nil {
		panic(err)
	}
	var puts func(string)
	purego.RegisterLibFunc(&puts, libc, "puts")
	puts("Calling C from Go without Cgo!")
}
