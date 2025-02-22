//go:build linux || windows

package miniaudio

import (
	"github.com/ebitengine/purego"
	"github.com/samborkent/miniaudio/internal/ma"
)

var (
	maDeviceConfigInit func(deviceConfig *ma.DeviceConfig, deviceType ma.DeviceType)
	maDeviceInit       func(context *ma.Context, config *ma.DeviceConfig, device *ma.Device) ma.Result
	maDeviceStart      func(device *ma.Device) ma.Result
	maDeviceUninit     func(device *ma.Device)
)

func init() {
	var err error

	lib, err := openLibrary()
	if err != nil {
		panic(err)
	}

	purego.RegisterLibFunc(&maDeviceConfigInit, lib, "ma_device_config_init")
	purego.RegisterLibFunc(&maDeviceInit, lib, "ma_device_init")
	purego.RegisterLibFunc(&maDeviceStart, lib, "ma_device_start")
	purego.RegisterLibFunc(&maDeviceUninit, lib, "ma_device_uninit")
}
