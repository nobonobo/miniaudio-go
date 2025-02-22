//go:build linux || windows

package miniaudio

import (
	"embed"
	"fmt"
	"runtime"

	"github.com/ebitengine/purego"
	"github.com/samborkent/miniaudio/internal/ma"
)

//go:embed build/*
var buildEmbed embed.FS

var (
	maDeviceConfigInit func(deviceConfig *ma.DeviceConfig, deviceType ma.DeviceType)
	maDeviceInit       func(context *ma.Context, config *ma.DeviceConfig, device *ma.Device) ma.Result
	maDeviceStart      func(device *ma.Device) ma.Result
	maDeviceUninit     func(device *ma.Device)
)

func init() {
	var err error

	lib, err := openLibrary(getSystemLibrary())
	if err != nil {
		panic(err)
	}

	purego.RegisterLibFunc(&maDeviceConfigInit, lib, "ma_device_config_init")
	purego.RegisterLibFunc(&maDeviceInit, lib, "ma_device_init")
	purego.RegisterLibFunc(&maDeviceStart, lib, "ma_device_start")
	purego.RegisterLibFunc(&maDeviceUninit, lib, "ma_device_uninit")
}

func getSystemLibrary() string {
	switch runtime.GOOS {
	case "linux":
		switch runtime.GOARCH {
		case "amd64":
			return "build/linux/libminiaudio-linux-amd64.so"
		case "arm64":
			return "build/linux/libminiaudio-linux-arm64.so"
		default:
			panic(fmt.Errorf("GOARCH=%s is not supported", runtime.GOARCH))
		}
	case "windows":
		switch runtime.GOARCH {
		case "amd64":
			return "build/windows/libminiaudio-windows-amd64.dll"
		case "arm64":
			return "build/windows/libminiaudio-windows-arm64.dll"
		default:
			panic(fmt.Errorf("GOARCH=%s is not supported", runtime.GOARCH))
		}
	default:
		panic(fmt.Errorf("GOOS=%s is not supported", runtime.GOOS))
	}
}
