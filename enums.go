package miniaudio

// ma_backend
type Backend int32

const (
	BackendWASAPI Backend = iota
	BackendDsound
	BackendWinMM
	BackendCoreAudio
	BackendSndio
	BackendAudio4
	BackendOSS
	BackendPulseAudio
	BackendALSA
	BackendJack
	BackendAAudio
	BackendOpenSL
	BackendWebAudio
	BackendCustom // Custom backend, with callbacks defined by the context config.
	BackendNull   // Must always be the last item. Lowest priority, and used as the terminator for backend enumeration.
	backendCount
)

const BackendCount = int32(backendCount)

// ma_channel_mix_mode
type ChannelMixMode int32

const (
	ChannelMixModeRectangular   ChannelMixMode = 0 // Simple averaging based on the plane(s) the channel is sitting on.
	ChannelMixModeSimple        ChannelMixMode = 1 // Drop excess channels; zeroed out extra channels.
	ChannelMixModeCustomWeights ChannelMixMode = 2 // Use custom weights specified in ChannelConverterConfig.
	ChannelMixModeDefault       ChannelMixMode = ChannelMixModeRectangular
)

// ma_device_state
type DeviceState int32

const (
	DeviceStateUninitialized DeviceState = iota
	DeviceStateStopped
	DeviceStateStarted
	DeviceStateStarting
	DeviceStateStopping
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
	FormatUnknown Format = iota // Mainly used for indicating an error, but also used as the default for the output format for decoders.
	FormatU8
	FormatS16 // Seems to be the most widely supported format.
	FormatS24 // Tightly packed. 3 bytes per sample.
	FormatS32
	FormatF32
	formatCount // Count of formats; ensures no gaps in the enum.
)

const FormatCount = int32(formatCount)

// ma_io_session_category
type IOSSessionCategory int32

const (
	IOSSessionCategoryDefault       IOSSessionCategory = 0
	IOSSessionCategoryNone          IOSSessionCategory = 1
	IOSSessionCategoryAmbient       IOSSessionCategory = 2
	IOSSessionCategorySoloAmbient   IOSSessionCategory = 3
	IOSSessionCategoryPlayback      IOSSessionCategory = 4
	IOSSessionCategoryRecord        IOSSessionCategory = 5
	IOSSessionCategoryPlayAndRecord IOSSessionCategory = 6
	IOSSessionCategoryMultiRoute    IOSSessionCategory = 7
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

// ma_standard_sample_rate
type StandardSampleRate int32

const (
	StandardSampleRate48000  StandardSampleRate = 48000 // Most common
	StandardSampleRate44100  StandardSampleRate = 44100
	StandardSampleRate32000  StandardSampleRate = 32000 // Lows
	StandardSampleRate24000  StandardSampleRate = 24000
	StandardSampleRate22050  StandardSampleRate = 22050
	StandardSampleRate88200  StandardSampleRate = 88200 // Highs
	StandardSampleRate96000  StandardSampleRate = 96000
	StandardSampleRate176400 StandardSampleRate = 176400
	StandardSampleRate192000 StandardSampleRate = 192000
	StandardSampleRate16000  StandardSampleRate = 16000 // Extreme lows
	StandardSampleRate11025  StandardSampleRate = 11025
	StandardSampleRate8000   StandardSampleRate = 8000
	StandardSampleRate352800 StandardSampleRate = 352800 // Extreme highs
	StandardSampleRate384000 StandardSampleRate = 384000

	StandardSampleRateMin   StandardSampleRate = StandardSampleRate8000   // Minimum value
	StandardSampleRateMax   StandardSampleRate = StandardSampleRate384000 // Maximum value
	StandardSampleRateCount int32              = 14                       // Total count of sample rates
)

// ma_thread_priority
type ThreadPriority int32

const (
	ThreadPriorityIdle     ThreadPriority = -5
	ThreadPriorityLowest   ThreadPriority = -4
	ThreadPriorityLow      ThreadPriority = -3
	ThreadPriorityNormal   ThreadPriority = -2
	ThreadPriorityHigh     ThreadPriority = -1
	ThreadPriorityHighest  ThreadPriority = 0
	ThreadPriorityRealtime ThreadPriority = 1
	ThreadPriorityDefault  ThreadPriority = 0
)

// ma_wasapi_usage
type WASAPIUsage int32

const (
	WASAPIUsageDefault  WASAPIUsage = 0 // Default usage.
	WASAPIUsageGames    WASAPIUsage = 1 // Usage optimized for games.
	WASAPIUsageProAudio WASAPIUsage = 2 // Usage optimized for professional audio.
)
