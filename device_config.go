package miniaudio

import (
	"fmt"
	"math/rand/v2"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/samborkent/miniaudio/internal/ma"
)

type DeviceConfig struct {
	DeviceType DeviceType
}

func (c DeviceConfig) toMA(config *ma.DeviceConfig) {
	config.DeviceType = c.DeviceType.toMA()
	config.SampleRate = 48000

	if c.DeviceType == DeviceTypePlayback || c.DeviceType == DeviceTypeDuplex {
		config.Playback.Format = ma.FormatF32
		config.Playback.Channels = 2
	}

	if c.DeviceType == DeviceTypeCapture || c.DeviceType == DeviceTypeDuplex {
		config.Capture.Format = ma.FormatF32
		config.Capture.Channels = 1
	}

	var dataCallback func(device *ma.Device, output, input ma.VoidPtr, frameCount uint32)

	switch c.DeviceType {
	case DeviceTypePlayback:
		dataCallback = func(_ *ma.Device, output, _ ma.VoidPtr, frameCount uint32) {
			samples := make([]float32, int(frameCount))

			for i := range frameCount {
				samples[i] = 2*rand.Float32() - 1
			}

			output = ma.VoidPtr(&samples[0])
		}
	case DeviceTypeCapture:
		dataCallback = func(_ *ma.Device, _, input ma.VoidPtr, frameCount uint32) {
			var floatType float32

			for i := range frameCount {
				fmt.Printf("%.3f", *(*float32)(unsafe.Pointer(uintptr(input) + uintptr(unsafe.Sizeof(floatType)*uintptr(i)))))
			}
		}
	case DeviceTypeDuplex:
		dataCallback = func(_ *ma.Device, output, input ma.VoidPtr, frameCount uint32) {
			var floatType float32

			for i := range frameCount {
				fmt.Printf("%.3f", *(*float32)(unsafe.Pointer(uintptr(input) + uintptr(unsafe.Sizeof(floatType)*uintptr(i)))))
			}

			output = input
		}
	default:
		panic("device type not supported")
	}

	config.DataCallback = ma.Proc(purego.NewCallback(dataCallback))
}
