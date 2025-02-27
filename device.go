package miniaudio

import (
	"sync/atomic"

	"github.com/samborkent/miniaudio/internal/ma"
)

type Device struct {
	Type DeviceType

	config *ma.DeviceConfig
	device *ma.Device

	initialized atomic.Bool
}

func NewDevice[T Formats](config DeviceConfig[T]) (*Device, error) {
	if !initialized.Load() {
		return nil, ErrNotInitialized
	}

	switch config.DeviceType {
	case DeviceTypePlayback, DeviceTypeCapture, DeviceTypeDuplex:
	default:
		return nil, ErrDeviceTypeUnknown
	}

	if config.DeviceType == DeviceTypePlayback ||
		config.DeviceType == DeviceTypeDuplex {
		if config.PlaybackCallback == nil {
			return nil, ErrNilPlaybackCallback
		}
	} else if config.DeviceType == DeviceTypeCapture ||
		config.DeviceType == DeviceTypeDuplex {
		if config.PlaybackCallback == nil {
			return nil, ErrNilCaptureCallback
		}
	}

	return &Device{
		Type:   config.DeviceType,
		config: config.toMA(),
	}, nil
}

func (d *Device) Init() error {
	if !initialized.Load() {
		return ErrNotInitialized
	}

	var device ma.Device

	result := maDeviceInit(nil, d.config, &device)
	if result != ma.Success {
		return convertResult(result)
	}

	d.device = &device
	d.initialized.Store(true)

	return nil
}

func (d *Device) PlaybackInfo() (DeviceInfo, error) {
	if !d.initialized.Load() {
		return DeviceInfo{}, ErrDeviceNotInitialized
	}

	var deviceInfo ma.DeviceInfo

	result := maDeviceGetInfo(d.device, ma.DeviceTypePlayback, &deviceInfo)
	if result != ma.Success {
		return DeviceInfo{}, convertResult(result)
	}

	return deviceInfoFromMA(deviceInfo), nil
}

func (d *Device) CaptureInfo() (DeviceInfo, error) {
	if !d.initialized.Load() {
		return DeviceInfo{}, ErrDeviceNotInitialized
	}

	var deviceInfo ma.DeviceInfo

	result := maDeviceGetInfo(d.device, ma.DeviceTypeCapture, &deviceInfo)
	if result != ma.Success {
		return DeviceInfo{}, convertResult(result)
	}

	return deviceInfoFromMA(deviceInfo), nil
}

func (d *Device) Start() error {
	if !d.initialized.Load() {
		return ErrDeviceNotInitialized
	}

	result := maDeviceStart(d.device)
	if result != ma.Success {
		return convertResult(result)
	}

	return nil
}

func (d *Device) Stop() error {
	if !d.initialized.Load() {
		return ErrDeviceNotInitialized
	}

	result := maDeviceStop(d.device)
	if result != ma.Success {
		return convertResult(result)
	}

	return nil
}

func (d *Device) Uninit() {
	if !d.initialized.Load() {
		return
	}

	maDeviceUninit(d.device)
}
