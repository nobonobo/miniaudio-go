package miniaudio

import "github.com/samborkent/miniaudio/internal/ma"

type DeviceType string

const (
	DeviceTypePlayback DeviceType = "playback"
	DeviceTypeCapture  DeviceType = "capture"
	DeviceTypeDuplex   DeviceType = "duplex"
)

func (t DeviceType) toMA() ma.DeviceType {
	switch t {
	case DeviceTypePlayback:
		return ma.DeviceTypePlayback
	case DeviceTypeCapture:
		return ma.DeviceTypeCapture
	case DeviceTypeDuplex:
		return ma.DeviceTypeDuplex
	default:
		panic("device type not supported")
	}
}
