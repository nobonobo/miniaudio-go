package miniaudio

import (
	"sync/atomic"

	"github.com/samborkent/miniaudio/internal/ma"
)

type Device struct {
	Config DeviceConfig

	config *ma.DeviceConfig
	device *ma.Device

	initialized atomic.Bool
}

func NewDevice(config DeviceConfig) (*Device, error) {
	if !initialized.Load() {
		return nil, ErrNotInitialized
	}

	return &Device{
		Config: config,
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

func (d *Device) Uninit() {
	if !d.initialized.Load() {
		return
	}

	maDeviceUninit(d.device)
}
