package miniaudio

// ma_device_type
type DeviceType int32

const (
	DeviceTypePlayback DeviceType = 1
	DeviceTypeCapture  DeviceType = 2
	DeviceTypeDuplex   DeviceType = DeviceTypePlayback | DeviceTypeCapture // 3
	DeviceTypeLoopback DeviceType = 4
)

// ma_format
type Format int32

const (
	FormatUnknown Format = 0 // Mainly used for indicating an error, but also used as the default for the output format for decoders.
	FormatU8      Format = 1
	FormatS16     Format = 2 // Seems to be the most widely supported format.
	FormatS24     Format = 3 // Tightly packed. 3 bytes per sample.
	FormatS32     Format = 4
	FormatF32     Format = 5
	FormatCount   Format = 6 // Count of formats; ensures no gaps in the enum.
)

// ma_performance_profile
type PerformanceProfile int32

const (
	PerformanceProfileLowLatency   PerformanceProfile = 0
	PerformanceProfileConservative PerformanceProfile = 1
)

// ma_resample_algorithm
type ResampleAlgorithm int32

const (
	ResampleAlgorithmLinear ResampleAlgorithm = 0 // Fastest, lowest quality. Optional low-pass filtering. Default.
	ResampleAlgorithmCustom ResampleAlgorithm = 1 // Custom resampling algorithm.
)
