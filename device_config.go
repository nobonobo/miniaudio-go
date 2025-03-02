package miniaudio

import (
	"sync/atomic"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/samborkent/miniaudio/internal/ma"
)

type DeviceConfig struct {
	DeviceType DeviceType
	SampleRate int
	Playback   FormatConfig
	Capture    FormatConfig

	dataCallback atomic.Uintptr
}

type FormatConfig struct {
	Format   Format
	Channels int
}

func (c *DeviceConfig) toMA() (*ma.DeviceConfig, error) {
	config := new(ma.DeviceConfig)

	config.DeviceType = c.DeviceType.toMA()
	config.SampleRate = uint32(c.SampleRate)
	config.Playback.Format = c.Playback.Format.toMA()

	playbackChannels := 0
	if c.Playback.Channels > 0 && c.Playback.Channels <= ma.MaxChannels {
		playbackChannels = c.Playback.Channels
	}

	config.Playback.Channels = uint32(playbackChannels)

	config.Capture.Format = c.Capture.Format.toMA()

	captureChannels := 0
	if c.Capture.Channels > 0 && c.Capture.Channels <= ma.MaxChannels {
		captureChannels = c.Capture.Channels
	}

	config.Capture.Channels = uint32(captureChannels)

	var dataCallback func(device *ma.Device, output, input unsafe.Pointer, frameCount uint32) uintptr
	var err error

	switch c.DeviceType {
	case DeviceTypePlayback:
		dataCallback, err = c.playbackCallback()
		if err != nil {
			return nil, err
		}
	case DeviceTypeCapture:
		dataCallback, err = c.captureCallback()
		if err != nil {
			return nil, err
		}
	case DeviceTypeDuplex:
		dataCallback, err = c.duplexCallback()
		if err != nil {
			return nil, err
		}
	default:
		return nil, ErrDeviceTypeNotSupported
	}

	config.DataCallback = ma.Proc(purego.NewCallback(dataCallback))

	return config, nil
}
