package miniaudio

import (
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/samborkent/miniaudio/internal/ma"
)

type (
	PlaybackCallback func(frameCount, channelCount int) [][]float32
	CaptureCallback  func(inputSamples []float32)
)

type DeviceConfig struct {
	DeviceType       DeviceType
	PlaybackCallback PlaybackCallback
	CaptureCallback  CaptureCallback
}

func (c DeviceConfig) toMA() *ma.DeviceConfig {
	config := new(ma.DeviceConfig)

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
			outputSamples := unsafe.Slice((*float32)(output), frameCount*config.Playback.Channels)

			gotSamples := c.PlaybackCallback(int(frameCount), int(config.Playback.Channels))

			for i := range frameCount {
				for c := range config.Playback.Channels {
					outputSamples[i*config.Playback.Channels+c] = gotSamples[i][c]
				}
			}

			return 0
		}
	case DeviceTypeCapture:
		dataCallback = func(_ *ma.Device, _, input unsafe.Pointer, frameCount uint32) uintptr {
			inputSamples := unsafe.Slice((*float32)(input), frameCount*config.Capture.Channels)

			c.CaptureCallback(inputSamples)

			return 0
		}
	case DeviceTypeDuplex:
		dataCallback = func(_ *ma.Device, output, input unsafe.Pointer, frameCount uint32) uintptr {
			inputSamples := unsafe.Slice((*float32)(input), frameCount*config.Capture.Channels)
			outputSamples := unsafe.Slice((*float32)(output), frameCount*config.Playback.Channels)

			go c.CaptureCallback(inputSamples)
			gotSamples := c.PlaybackCallback(int(frameCount), int(config.Playback.Channels))

			for i := range frameCount {
				for c := range config.Playback.Channels {
					outputSamples[i*config.Playback.Channels+c] = gotSamples[i][c]
				}
			}

			return 0
		}
	default:
		panic("device type not supported")
	}

	config.DataCallback = ma.Proc(purego.NewCallback(dataCallback))

	return config
}
