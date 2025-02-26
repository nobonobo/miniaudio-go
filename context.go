package miniaudio

import (
	"sync/atomic"
	"unsafe"

	"github.com/samborkent/miniaudio/cutil"
	"github.com/samborkent/miniaudio/internal/ma"
)

type Context struct {
	context *ma.Context

	initialized atomic.Bool
}

func NewContext() *Context {
	return &Context{}
}

func (c *Context) Init() error {
	if !initialized.Load() {
		return ErrNotInitialized
	}

	var context ma.Context

	result := maContextInit(nil, 0, nil, &context)
	if result != ma.Success {
		return convertResult(result)
	}

	c.context = &context
	c.initialized.Store(true)

	return nil
}

func (c *Context) GetDevices() (playbackDevices, captureDevices []DeviceInfo, err error) {
	pDevices := &ma.DeviceInfo{}
	cDevices := &ma.DeviceInfo{}

	var pCount, cCount uint32

	result := maContextGetDevices(c.context, &pDevices, &pCount, &cDevices, &cCount)
	if result != ma.Success {
		return nil, nil, convertResult(result)
	}

	maPlaybackDevices := unsafe.Slice(pDevices, pCount)
	playbackDevices = make([]DeviceInfo, len(maPlaybackDevices))

	for i, deviceInfo := range maPlaybackDevices {
		var isDefault bool

		if deviceInfo.IsDefault > 0 {
			isDefault = true
		}

		playbackDevices[i] = DeviceInfo{
			ID:          cutil.String(deviceInfo.ID[:]),
			Name:        cutil.String(deviceInfo.Name[:]),
			IsDefault:   isDefault,
			DataFormats: make([]DataFormat, deviceInfo.NativeDataFormatCount),
		}

		for j := range deviceInfo.NativeDataFormatCount {
			playbackDevices[i].DataFormats[j] = DataFormat{
				Format:     formatFromMA(deviceInfo.NativeDataFormats[i].Format),
				Channels:   int(deviceInfo.NativeDataFormats[i].Channels),
				SampleRate: int(deviceInfo.NativeDataFormats[i].SampleRate),
				Flags:      deviceInfo.NativeDataFormats[i].Flags,
			}
		}
	}

	maCapturePlaybackDevices := unsafe.Slice(cDevices, cCount)
	captureDevices = make([]DeviceInfo, len(maCapturePlaybackDevices))

	for i, deviceInfo := range maCapturePlaybackDevices {
		var isDefault bool

		if deviceInfo.IsDefault > 0 {
			isDefault = true
		}

		captureDevices[i] = DeviceInfo{
			ID:          cutil.String(deviceInfo.ID[:]),
			Name:        cutil.String(deviceInfo.Name[:]),
			IsDefault:   isDefault,
			DataFormats: make([]DataFormat, deviceInfo.NativeDataFormatCount),
		}

		for j := range deviceInfo.NativeDataFormatCount {
			playbackDevices[i].DataFormats[j] = DataFormat{
				Format:     formatFromMA(deviceInfo.NativeDataFormats[i].Format),
				Channels:   int(deviceInfo.NativeDataFormats[i].Channels),
				SampleRate: int(deviceInfo.NativeDataFormats[i].SampleRate),
				Flags:      deviceInfo.NativeDataFormats[i].Flags,
			}
		}
	}

	return playbackDevices, captureDevices, nil
}
