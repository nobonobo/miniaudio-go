package miniaudio

import (
	"github.com/samborkent/miniaudio/internal/ma"
)

type Device struct {
	Config DeviceConfig

	config *ma.DeviceConfig
	device *ma.Device
}

func NewDevice(config DeviceConfig) *Device {
	var maConfig ma.DeviceConfig
	config.toMA(&maConfig)

	return &Device{
		Config: config,
		config: &maConfig,
	}
}

func (d *Device) Init() error {
	var device ma.Device

	result := maDeviceInit(nil, d.config, &device)
	if result != ma.Success {
		return convertResult(result)
	}

	d.device = &device

	return nil
}

func (d *Device) Start() error {
	result := maDeviceStart(d.device)
	if result != ma.Success {
		return convertResult(result)
	}

	return nil
}

func (d *Device) Uninit() {
	maDeviceUninit(d.device)
}
