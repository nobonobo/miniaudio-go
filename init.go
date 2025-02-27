//go:build linux || windows

package miniaudio

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/ebitengine/purego"
	"github.com/samborkent/miniaudio/internal/ma"
)

var (
	maContextInit       func(backends []ma.Backend, backendCount uint32, config *ma.ContextConfig, context *ma.Context) ma.Result
	maContextGetDevices func(context *ma.Context, playbackDevices **ma.DeviceInfo, playbackDeviceCount *uint32, captureDevices **ma.DeviceInfo, captureDeviceCount *uint32) ma.Result
	maContextUninit     func(context *ma.Context)

	maDeviceConfigInit func(deviceConfig *ma.DeviceConfig, deviceType ma.DeviceType)
	maDeviceInit       func(context *ma.Context, config *ma.DeviceConfig, device *ma.Device) ma.Result
	maDeviceGetInfo    func(device *ma.Device, deviceType ma.DeviceType, deviceInfo *ma.DeviceInfo) ma.Result
	maDeviceStart      func(device *ma.Device) ma.Result
	maDeviceUninit     func(device *ma.Device)
)

var (
	initOnce    sync.Once
	initialized atomic.Bool
	initErr     atomic.Value
)

// Init loads in the miniaudio library and registers the relevant functions.
func Init() error {
	if !initialized.Load() {
		initOnce.Do(func() {
			lib, err := openLibrary()
			if err != nil {
				initErr.Store(fmt.Errorf("miniaudio: opening library: %w", err))
				return
			}

			purego.RegisterLibFunc(&maContextInit, lib, "ma_context_init")
			purego.RegisterLibFunc(&maContextGetDevices, lib, "ma_context_get_devices")
			purego.RegisterLibFunc(&maContextUninit, lib, "ma_context_uninit")
			purego.RegisterLibFunc(&maDeviceConfigInit, lib, "ma_device_config_init")
			purego.RegisterLibFunc(&maDeviceInit, lib, "ma_device_init")
			purego.RegisterLibFunc(&maDeviceGetInfo, lib, "ma_device_get_info")
			purego.RegisterLibFunc(&maDeviceStart, lib, "ma_device_start")
			purego.RegisterLibFunc(&maDeviceUninit, lib, "ma_device_uninit")

			initialized.Store(true)
		})
	}

	errAny := initErr.Load()
	if errAny != nil {
		err, ok := errAny.(error)
		if ok {
			return err
		}
	}

	return nil
}
