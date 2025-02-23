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

	var dataCallback func(device *ma.Device, output, input unsafe.Pointer, frameCount uint32) uintptr

	switch c.DeviceType {
	case DeviceTypePlayback:
		dataCallback = func(_ *ma.Device, output, _ unsafe.Pointer, frameCount uint32) uintptr {
			samples := unsafe.Slice((*float32)(output), frameCount*config.Playback.Channels)

			for i := range frameCount {
				sample := 2*rand.Float32() - 1

				for c := range config.Playback.Channels {
					samples[i*config.Playback.Channels+c] = sample
				}
			}

			return 0
		}
	case DeviceTypeCapture:
		dataCallback = func(_ *ma.Device, _, input unsafe.Pointer, frameCount uint32) uintptr {
			var floatType float32

			for i := range frameCount {
				fmt.Printf("%.3f", *(*float32)(unsafe.Pointer(uintptr(input) + uintptr(unsafe.Sizeof(floatType)*uintptr(i)))))
			}

			return 0
		}
	case DeviceTypeDuplex:
		dataCallback = func(_ *ma.Device, output, input unsafe.Pointer, frameCount uint32) uintptr {
			inputSamples := unsafe.Slice((*float32)(input), frameCount)
			outputSamples := unsafe.Slice((*float32)(output), frameCount*config.Playback.Channels)

			for i := range frameCount {
				for c := range config.Playback.Channels {
					outputSamples[i*config.Playback.Channels+c] = inputSamples[i]
				}
			}

			return 0
		}
	default:
		panic("device type not supported")
	}

	config.DataCallback = ma.Proc(purego.NewCallback(dataCallback))
}
