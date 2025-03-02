package miniaudio

import (
	"log/slog"
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

	if (d.Config.DeviceType == DeviceTypePlayback ||
		d.Config.DeviceType == DeviceTypeDuplex) &&
		d.config.Playback.Channels == 0 {
		playbackInfo, err := d.PlaybackInfo()
		if err != nil {
			return err
		}

		if len(playbackInfo.DataFormats) == 0 {
			return ErrChannelsUnknown
		}

		if len(playbackInfo.DataFormats) > 1 {
			slog.Warn("multiple playback formats detected!")
		}

		d.Config.Playback.Channels = playbackInfo.DataFormats[0].Channels
	}

	if (d.Config.DeviceType == DeviceTypeCapture ||
		d.Config.DeviceType == DeviceTypeDuplex) &&
		d.config.Capture.Channels == 0 {
		captureInfo, err := d.CaptureInfo()
		if err != nil {
			return err
		}

		if len(captureInfo.DataFormats) == 0 {
			return ErrChannelsUnknown
		}

		if len(captureInfo.DataFormats) > 1 {
			slog.Warn("multiple capture formats detected!")
		}

		d.Config.Capture.Channels = captureInfo.DataFormats[0].Channels
	}

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
