package miniaudio

import (
	"sync/atomic"
	"unsafe"

	"github.com/samborkent/miniaudio/internal/cutil"
	"github.com/samborkent/miniaudio/internal/ma"
)

type Context struct {
	Backend Backend

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

	backend, err := backendFromMA(context.Backend)
	if err != nil {
		return err
	}

	c.Backend = backend
	c.context = &context
	c.initialized.Store(true)

	return nil
}

func (c *Context) GetDevices() (playbackDevices, captureDevices []DeviceInfo, err error) {
	if !c.initialized.Load() {
		return nil, nil, ErrContextNotInitialized
	}

	pDevices := &ma.DeviceInfo{}
	cDevices := &ma.DeviceInfo{}

	var pCount, cCount uint32

	result := maContextGetDevices(c.context, &pDevices, &pCount, &cDevices, &cCount)
	if result != ma.Success {
		return nil, nil, convertResult(result)
	}

	maPlaybackDevices := unsafe.Slice(pDevices, pCount)
	playbackDevices = make([]DeviceInfo, len(maPlaybackDevices))

	for i := range pCount {
		var isDefault bool

		if maPlaybackDevices[i].IsDefault > 0 {
			isDefault = true
		}

		playbackDevices[i] = DeviceInfo{
			ID:          cutil.String(maPlaybackDevices[i].ID[:]),
			Name:        cutil.String(maPlaybackDevices[i].Name[:]),
			IsDefault:   isDefault,
			DataFormats: make([]DataFormat, maPlaybackDevices[i].NativeDataFormatCount),
		}

		for j := range maPlaybackDevices[i].NativeDataFormatCount {
			playbackDevices[i].DataFormats[j] = DataFormat{
				Format:     formatFromMA(maPlaybackDevices[i].NativeDataFormats[j].Format),
				Channels:   int(maPlaybackDevices[i].NativeDataFormats[j].Channels),
				SampleRate: int(maPlaybackDevices[i].NativeDataFormats[j].SampleRate),
				Flags:      maPlaybackDevices[i].NativeDataFormats[j].Flags,
			}
		}
	}

	maCapturePlaybackDevices := unsafe.Slice(cDevices, cCount)
	captureDevices = make([]DeviceInfo, len(maCapturePlaybackDevices))

	for i := range cCount {
		var isDefault bool

		if maCapturePlaybackDevices[i].IsDefault > 0 {
			isDefault = true
		}

		captureDevices[i] = DeviceInfo{
			ID:          cutil.String(maCapturePlaybackDevices[i].ID[:]),
			Name:        cutil.String(maCapturePlaybackDevices[i].Name[:]),
			IsDefault:   isDefault,
			DataFormats: make([]DataFormat, maCapturePlaybackDevices[i].NativeDataFormatCount),
		}

		for j := range maCapturePlaybackDevices[i].NativeDataFormatCount {
			playbackDevices[i].DataFormats[j] = DataFormat{
				Format:     formatFromMA(maCapturePlaybackDevices[i].NativeDataFormats[j].Format),
				Channels:   int(maCapturePlaybackDevices[i].NativeDataFormats[j].Channels),
				SampleRate: int(maCapturePlaybackDevices[i].NativeDataFormats[j].SampleRate),
				Flags:      maCapturePlaybackDevices[i].NativeDataFormats[j].Flags,
			}
		}
	}

	return playbackDevices, captureDevices, nil
}

func (c *Context) Uninit() {
	if !c.initialized.Load() {
		return
	}

	maContextUninit(c.context)
}
