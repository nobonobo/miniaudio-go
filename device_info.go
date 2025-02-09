package miniaudio

const NativeDataFormatsCount = 64 // FormatCount * StandardSampleRateCount * MaxChannels

type DeviceInfo struct {
	ID                    DeviceID                      // ma_device_id
	Name                  [MaxDeviceNameLength + 1]byte // char name[MA_MAX_DEVICE_NAME_LENGTH + 1]
	IsDefault             uint32                        // ma_bool32
	NativeDataFormatCount uint32                        // ma_uint32
	NativeDataFormats     [NativeDataFormatsCount]struct {
		Format     Format // ma_format
		Channels   uint32 // ma_uint32
		SampleRate uint32 // ma_uint32
		Flags      uint32 // ma_uint32 (combination of MA_DATA_FORMAT_FLAG_* flags)
	}
}
