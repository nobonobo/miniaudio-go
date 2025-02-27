package miniaudio

import (
	"github.com/samborkent/miniaudio/cutil"
	"github.com/samborkent/miniaudio/internal/ma"
)

// ID                    DeviceID                      // ma_device_id
// Name                  [MaxDeviceNameLength + 1]byte // char name[MA_MAX_DEVICE_NAME_LENGTH + 1]
// IsDefault             uint32                        // ma_bool32
// NativeDataFormatCount uint32                        // ma_uint32
// NativeDataFormats     [NativeDataFormatsCount]struct {
// 	Format     Format // ma_format
// 	Channels   uint32 // ma_uint32
// 	SampleRate uint32 // ma_uint32
// 	Flags      uint32 // ma_uint32 (combination of MA_DATA_FORMAT_FLAG_* flags)
// }

type DeviceInfo struct {
	ID          string
	Name        string
	IsDefault   bool
	DataFormats []DataFormat
}

type DataFormat struct {
	Format     Format
	Channels   int
	SampleRate int
	Flags      uint32
}

func deviceInfoFromMA(deviceInfo ma.DeviceInfo) DeviceInfo {
	var isDefault bool

	if deviceInfo.IsDefault > 0 {
		isDefault = true
	}

	di := DeviceInfo{
		ID:          cutil.String(deviceInfo.ID[:]),
		Name:        cutil.String(deviceInfo.Name[:]),
		IsDefault:   isDefault,
		DataFormats: make([]DataFormat, deviceInfo.NativeDataFormatCount),
	}

	for i := range deviceInfo.NativeDataFormatCount {
		di.DataFormats[i] = DataFormat{
			Format:     formatFromMA(deviceInfo.NativeDataFormats[i].Format),
			Channels:   int(deviceInfo.NativeDataFormats[i].Channels),
			SampleRate: int(deviceInfo.NativeDataFormats[i].SampleRate),
			Flags:      deviceInfo.NativeDataFormats[i].Flags,
		}
	}

	return di
}
