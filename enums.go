package miniaudio

// ma_channel_mix_mode
type ChannelMixMode int32

const (
	ChannelMixModeRectangular   ChannelMixMode = 0 // Simple averaging based on the plane(s) the channel is sitting on.
	ChannelMixModeSimple        ChannelMixMode = 1 // Drop excess channels; zeroed out extra channels.
	ChannelMixModeCustomWeights ChannelMixMode = 2 // Use custom weights specified in ChannelConverterConfig.
	ChannelMixModeDefault       ChannelMixMode = ChannelMixModeRectangular
)

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

// ma_share_mode
type ShareMode int32

const (
	ShareModeShared    ShareMode = 0 // Shared mode.
	ShareModeExclusive ShareMode = 1 // Exclusive mode.
)

// ma_wasapi_usage
type WASAPIUsage int32

const (
	WASAPIUsageDefault  WASAPIUsage = 0 // Default usage.
	WASAPIUsageGames    WASAPIUsage = 1 // Usage optimized for games.
	WASAPIUsageProAudio WASAPIUsage = 2 // Usage optimized for professional audio.
)
