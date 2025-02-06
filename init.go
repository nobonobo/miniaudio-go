//go:build linux || windows

package miniaudio

import (
	"fmt"
	"runtime"
)

var lib uintptr

func init() {
	var err error

	lib, err = openLibrary(getSystemLibrary())
	if err != nil {
		panic(err)
	}

	// purego.RegisterLibFunc(&maEngineInit, lib, "ma_engine_init")

	// var engine maEngine
	// maEngineInit(nil, &engine)
}

func getSystemLibrary() string {
	switch runtime.GOOS {
	case "linux":
		switch runtime.GOARCH {
		case "amd64":
			return "build/libminiaudio-linux-amd64.so"
		case "arm64":
			return "build/libminiaudio-linux-arm64.so"
		default:
			panic(fmt.Errorf("GOARCH=%s is not supported", runtime.GOARCH))
		}
	case "windows":
		switch runtime.GOARCH {
		case "amd64":
			return "build/libminiaudio-windows-amd64.dll"
		case "arm64":
			return "build/libminiaudio-windows-arm64.dll"
		default:
			panic(fmt.Errorf("GOARCH=%s is not supported", runtime.GOARCH))
		}
	default:
		panic(fmt.Errorf("GOOS=%s is not supported", runtime.GOOS))
	}
}
