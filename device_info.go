package miniaudio

import (
	"github.com/samborkent/miniaudio/internal/cutil"
	"github.com/samborkent/miniaudio/internal/ma"
)

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

func deviceIDToMA(deviceID string) ma.DeviceID {
	var maDeviceID ma.DeviceID

	idBytes := []byte(deviceID)

	if len(idBytes) > 254 {
		idBytes = idBytes[:254]
	}

	idBytes = append(idBytes, '\x00')

	copy(maDeviceID[:], idBytes)

	return maDeviceID
}
