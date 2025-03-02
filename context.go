package miniaudio

import (
	"sync/atomic"
	"unsafe"

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
		playbackDevices[i] = deviceInfoFromMA(maPlaybackDevices[i])
	}

	maCaptureDevices := unsafe.Slice(cDevices, cCount)
	captureDevices = make([]DeviceInfo, len(maCaptureDevices))

	for i := range cCount {
		captureDevices[i] = deviceInfoFromMA(maCaptureDevices[i])
	}

	return playbackDevices, captureDevices, nil
}

func (c *Context) GetDefaultPlayback() (DeviceInfo, error) {
	if !c.initialized.Load() {
		return DeviceInfo{}, ErrContextNotInitialized
	}

	pDevices := &ma.DeviceInfo{}

	var pCount uint32

	result := maContextGetDevices(c.context, &pDevices, &pCount, nil, nil)
	if result != ma.Success {
		return DeviceInfo{}, convertResult(result)
	}

	maPlaybackDevices := unsafe.Slice(pDevices, pCount)

	for i := range pCount {
		if maPlaybackDevices[i].IsDefault > 0 {
			return deviceInfoFromMA(maPlaybackDevices[i]), nil
		}
	}

	return DeviceInfo{}, ErrNoDefaultDevice
}

func (c *Context) GetDefaultCapture() (DeviceInfo, error) {
	if !c.initialized.Load() {
		return DeviceInfo{}, ErrContextNotInitialized
	}

	cDevices := &ma.DeviceInfo{}

	var cCount uint32

	result := maContextGetDevices(c.context, nil, nil, &cDevices, &cCount)
	if result != ma.Success {
		return DeviceInfo{}, convertResult(result)
	}

	maCaptureDevices := unsafe.Slice(cDevices, cCount)

	for i := range cCount {
		if maCaptureDevices[i].IsDefault > 0 {
			return deviceInfoFromMA(maCaptureDevices[i]), nil
		}
	}

	return DeviceInfo{}, ErrNoDefaultDevice
}

func (c *Context) GetDeviceInfo(deviceType DeviceType, deviceID string) (DeviceInfo, error) {
	if !c.initialized.Load() {
		return DeviceInfo{}, ErrContextNotInitialized
	}

	var deviceInfo ma.DeviceInfo

	maDeviceID := deviceIDToMA(deviceID)

	result := maContextGetDeviceInfo(c.context, deviceType.toMA(), &maDeviceID, &deviceInfo)
	if result != ma.Success {
		return DeviceInfo{}, convertResult(result)
	}

	return deviceInfoFromMA(deviceInfo), nil
}

func (c *Context) Uninit() {
	if !c.initialized.Load() {
		return
	}

	maContextUninit(c.context)
}
