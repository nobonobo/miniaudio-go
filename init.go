//go:build linux || windows

package miniaudio

import (
	"embed"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/ebitengine/purego"
	"github.com/samborkent/miniaudio/internal/ma"
)

//go:embed build/*
var buildEmbed embed.FS

const rootDir = "build"

var (
	maContextInit          func(backends []ma.Backend, backendCount uint32, config *ma.ContextConfig, context *ma.Context) ma.Result
	maContextGetDevices    func(context *ma.Context, playbackDevices **ma.DeviceInfo, playbackDeviceCount *uint32, captureDevices **ma.DeviceInfo, captureDeviceCount *uint32) ma.Result
	maContextGetDeviceInfo func(context *ma.Context, deviceType ma.DeviceType, deviceID *ma.DeviceID, deviceInfo *ma.DeviceInfo) ma.Result
	maContextUninit        func(context *ma.Context)

	maDeviceConfigInit func(deviceConfig *ma.DeviceConfig, deviceType ma.DeviceType)
	maDeviceInit       func(context *ma.Context, config *ma.DeviceConfig, device *ma.Device) ma.Result
	maDeviceGetContext func(device *ma.Device) *ma.Context
	maDeviceGetInfo    func(device *ma.Device, deviceType ma.DeviceType, deviceInfo *ma.DeviceInfo) ma.Result
	maDeviceStart      func(device *ma.Device) ma.Result
	maDeviceStop       func(device *ma.Device) ma.Result
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
			dirEntries, err := buildEmbed.ReadDir(rootDir)
			if err != nil {
				initErr.Store(fmt.Errorf("reading embedded directory: %w", err))
				return
			}

			var lib uintptr

			for _, dirEntry := range dirEntries {
				if dirEntry.IsDir() {
					continue
				}

				if !strings.Contains(dirEntry.Name(), runtime.GOOS) {
					continue
				}

				if !strings.Contains(dirEntry.Name(), runtime.GOARCH) {
					continue
				}

				tmpFile, err := os.CreateTemp("", dirEntry.Name())
				if err != nil {
					initErr.Store(fmt.Errorf("creating temporary file: %w", err))
					return
				}

				defer func() {
					if err := os.Remove(tmpFile.Name()); err != nil {
						initErr.Store(fmt.Errorf("removing temporary file: %w", err))
						return
					}
				}()

				libFile, err := buildEmbed.Open(path.Join(rootDir, dirEntry.Name()))
				if err != nil {
					initErr.Store(fmt.Errorf("opening embedded library file: %w", err))
					return
				}

				_, err = io.Copy(tmpFile, libFile)
				if err != nil {
					initErr.Store(fmt.Errorf("writing library contents to temporary file: %w", err))
					return
				}

				lib, err = openLibrary(tmpFile.Name())
				if err != nil {
					initErr.Store(fmt.Errorf("loading shared library: %w", err))
					return
				}

				break
			}

			if lib == 0 {
				initErr.Store(ErrLibraryNotFound)
				return
			}

			purego.RegisterLibFunc(&maContextInit, lib, "ma_context_init")
			purego.RegisterLibFunc(&maContextGetDevices, lib, "ma_context_get_devices")
			purego.RegisterLibFunc(&maContextGetDeviceInfo, lib, "ma_context_get_device_info")
			purego.RegisterLibFunc(&maContextUninit, lib, "ma_context_uninit")
			purego.RegisterLibFunc(&maDeviceConfigInit, lib, "ma_device_config_init")
			purego.RegisterLibFunc(&maDeviceInit, lib, "ma_device_init")
			purego.RegisterLibFunc(&maDeviceGetContext, lib, "ma_device_get_context")
			purego.RegisterLibFunc(&maDeviceGetInfo, lib, "ma_device_get_info")
			purego.RegisterLibFunc(&maDeviceStart, lib, "ma_device_start")
			purego.RegisterLibFunc(&maDeviceStop, lib, "ma_device_stop")
			purego.RegisterLibFunc(&maDeviceUninit, lib, "ma_device_uninit")

			initialized.Store(true)
		})
	}

	errAny := initErr.Load()
	if errAny != nil {
		err, ok := errAny.(error)
		if ok {
			return fmt.Errorf("miniaudio: %w", err)
		}
	}

	return nil
}
