package miniaudio

import (
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/samborkent/miniaudio/internal/ma"
)

type (
	PlaybackCallback[T Formats] func(frameCount, channelCount int) [][]T
	CaptureCallback[T Formats]  func(inputSamples []T, frameCount, channelCount int)
)

type DeviceConfig[T Formats] struct {
	DeviceType       DeviceType
	PlaybackCallback PlaybackCallback[T]
	CaptureCallback  CaptureCallback[T]
}

func (c DeviceConfig[T]) toMA() *ma.DeviceConfig {
	config := new(ma.DeviceConfig)

	config.DeviceType = c.DeviceType.toMA()
	config.SampleRate = 48000

	switch any(*new(T)).(type) {
	case uint8:
		config.Playback.Format = ma.FormatU8
		config.Capture.Format = ma.FormatU8
	case int16:
		config.Playback.Format = ma.FormatS16
		config.Capture.Format = ma.FormatS16
	case int32:
		config.Playback.Format = ma.FormatS32
		config.Capture.Format = ma.FormatS32
	case float32:
		config.Playback.Format = ma.FormatF32
		config.Capture.Format = ma.FormatF32
	}

	if c.DeviceType == DeviceTypePlayback || c.DeviceType == DeviceTypeDuplex {
		config.Playback.Channels = 2
	}

	if c.DeviceType == DeviceTypeCapture || c.DeviceType == DeviceTypeDuplex {
		config.Capture.Channels = 1
	}

	var dataCallback func(device *ma.Device, output, input unsafe.Pointer, frameCount uint32) uintptr

	switch c.DeviceType {
	case DeviceTypePlayback:
		dataCallback = func(_ *ma.Device, output, _ unsafe.Pointer, frameCount uint32) uintptr {
			outputSamples := unsafe.Slice((*T)(output), frameCount*config.Playback.Channels)

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
			inputSamples := unsafe.Slice((*T)(input), frameCount*config.Capture.Channels)

			c.CaptureCallback(inputSamples, int(frameCount), int(config.Capture.Channels))

			return 0
		}
	case DeviceTypeDuplex:
		dataCallback = func(_ *ma.Device, output, input unsafe.Pointer, frameCount uint32) uintptr {
			inputSamples := unsafe.Slice((*T)(input), frameCount*config.Capture.Channels)
			outputSamples := unsafe.Slice((*T)(output), frameCount*config.Playback.Channels)

			go c.CaptureCallback(inputSamples, int(frameCount), int(config.Playback.Channels))
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
