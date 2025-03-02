package miniaudio

import (
	"math"
	"sync/atomic"

	"github.com/samborkent/miniaudio/internal/ma"
)

type Device struct {
	Config *DeviceConfig

	config *ma.DeviceConfig
	device *ma.Device

	initialized atomic.Bool
}

func NewDevice(config *DeviceConfig) (*Device, error) {
	if !initialized.Load() {
		return nil, ErrNotInitialized
	}

	switch config.DeviceType {
	case DeviceTypePlayback, DeviceTypeCapture, DeviceTypeDuplex:
	default:
		return nil, ErrDeviceTypeUnknown
	}

	if config.SampleRate < 0 || config.SampleRate > math.MaxUint32 {
		return nil, ErrSampleRateInvalid
	}

	if config.Playback.Channels < 0 || config.Playback.Channels > ma.MaxChannels ||
		config.Capture.Channels < 0 || config.Capture.Channels > ma.MaxChannels {
		return nil, ErrChannelsInvalid
	}

	if config.dataCallback.Load() == 0 {
		return nil, ErrNilCallback
	}

	cfg, err := config.toMA()
	if err != nil {
		return nil, err
	}

	return &Device{
		Config: config,
		config: cfg,
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

func (d *Device) InitWithContext(context *Context) error {
	if !initialized.Load() {
		return ErrNotInitialized
	}

	var device ma.Device

	result := maDeviceInit(context.context, d.config, &device)
	if result != ma.Success {
		return convertResult(result)
	}

	d.device = &device
	d.initialized.Store(true)

	return nil
}

func (d *Device) PlaybackInfo() (DeviceInfo, error) {
	return d.deviceInfo(DeviceTypePlayback)
}

func (d *Device) CaptureInfo() (DeviceInfo, error) {
	return d.deviceInfo(DeviceTypeCapture)
}

func (d *Device) deviceInfo(deviceType DeviceType) (DeviceInfo, error) {
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
