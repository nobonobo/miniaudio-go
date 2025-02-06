package miniaudio

import "unsafe"

type Channel uint8

// ma_device_config
type DeviceConfig struct {
	DeviceType                DeviceType         // ma_device_type
	SampleRate                uint32             // ma_uint32
	PeriodSizeInFrames        uint32             // ma_uint32
	PeriodSizeInMilliseconds  uint32             // ma_uint32
	Periods                   uint32             // ma_uint32
	PerformanceProfile        PerformanceProfile // ma_performance_profile
	NoPreSilencedOutputBuffer bool               // ma_bool8
	NoClip                    bool               // ma_bool8
	NoDisableDenormals        bool               // ma_bool8
	NoFixedSizedCallback      bool               // ma_bool8
	DataCallback              uintptr            // ma_device_data_proc (function pointer)
	NotificationCallback      uintptr            // ma_device_notification_proc (function pointer)
	StopCallback              uintptr            // ma_stop_proc (function pointer)
	PUserData                 unsafe.Pointer     // void*
	Resampling                ResamplerConfig    // ma_resampler_config

	Playback struct {
		PDeviceID   *DeviceID
		Format      Format   // ma_format
		Channels    uint32   // ma_uint32
		PChannelMap *Channel // ma_channel*
		//
		// TODO: left off here!!!
		//
		ChannelMixMode                  int32 // ma_channel_mix_mode (assumed int32)
		CalculateLFEFromSpatialChannels bool  // ma_bool32
		ShareMode                       int32 // ma_share_mode (assumed int32)
	}

	Capture struct {
		PDeviceID                       *DeviceID
		Format                          Format   // ma_format
		Channels                        uint32   // ma_uint32
		PChannelMap                     *Channel // *ma_channel
		ChannelMixMode                  int32    // ma_channel_mix_mode (assumed int32)
		CalculateLFEFromSpatialChannels bool     // ma_bool32
		ShareMode                       int32    // ma_share_mode (assumed int32)
	}

	WASAPI struct {
		Usage                  int32  // ma_wasapi_usage (assumed int32)
		NoAutoConvertSRC       bool   // ma_bool8
		NoDefaultQualitySRC    bool   // ma_bool8
		NoAutoStreamRouting    bool   // ma_bool8
		NoHardwareOffloading   bool   // ma_bool8
		LoopbackProcessID      uint32 // ma_uint32
		LoopbackProcessExclude bool   // ma_bool8
	}

	ALSA struct {
		NoMMap         bool // ma_bool32
		NoAutoFormat   bool // ma_bool32
		NoAutoChannels bool // ma_bool32
		NoAutoResample bool // ma_bool32
	}

	Pulse struct {
		PStreamNamePlayback *string
		PStreamNameCapture  *string
	}

	CoreAudio struct {
		AllowNominalSampleRateChange bool // ma_bool32
	}

	OpenSL struct {
		StreamType                     int32 // ma_opensl_stream_type (assumed int32)
		RecordingPreset                int32 // ma_opensl_recording_preset (assumed int32)
		EnableCompatibilityWorkarounds bool  // ma_bool32
	}

	AAudio struct {
		Usage                          int32 // ma_aaudio_usage (assumed int32)
		ContentType                    int32 // ma_aaudio_content_type (assumed int32)
		InputPreset                    int32 // ma_aaudio_input_preset (assumed int32)
		AllowedCapturePolicy           int32 // ma_aaudio_allowed_capture_policy (assumed int32)
		NoAutoStartAfterReroute        bool  // ma_bool32
		EnableCompatibilityWorkarounds bool  // ma_bool32
	}
}

// struct ma_device_config
// {
//     ma_device_type deviceType;
//     ma_uint32 sampleRate;
//     ma_uint32 periodSizeInFrames;
//     ma_uint32 periodSizeInMilliseconds;
//     ma_uint32 periods;
//     ma_performance_profile performanceProfile;
//     ma_bool8 noPreSilencedOutputBuffer; /* When set to true, the contents of the output buffer passed into the data callback will be left undefined rather than initialized to silence. */
//     ma_bool8 noClip;                    /* When set to true, the contents of the output buffer passed into the data callback will not be clipped after returning. Only applies when the playback sample format is f32. */
//     ma_bool8 noDisableDenormals;        /* Do not disable denormals when firing the data callback. */
//     ma_bool8 noFixedSizedCallback;      /* Disables strict fixed-sized data callbacks. Setting this to true will result in the period size being treated only as a hint to the backend. This is an optimization for those who don't need fixed sized callbacks. */
//     ma_device_data_proc dataCallback;
//     ma_device_notification_proc notificationCallback;
//     ma_stop_proc stopCallback;
//     void* pUserData;
//     ma_resampler_config resampling;
//     struct
//     {
//         const ma_device_id* pDeviceID;
//         ma_format format;
//         ma_uint32 channels;
//         ma_channel* pChannelMap;
//         ma_channel_mix_mode channelMixMode;
//         ma_bool32 calculateLFEFromSpatialChannels;  /* When an output LFE channel is present, but no input LFE, set to true to set the output LFE to the average of all spatial channels (LR, FR, etc.). Ignored when an input LFE is present. */
//         ma_share_mode shareMode;
//     } playback;
//     struct
//     {
//         const ma_device_id* pDeviceID;
//         ma_format format;
//         ma_uint32 channels;
//         ma_channel* pChannelMap;
//         ma_channel_mix_mode channelMixMode;
//         ma_bool32 calculateLFEFromSpatialChannels;  /* When an output LFE channel is present, but no input LFE, set to true to set the output LFE to the average of all spatial channels (LR, FR, etc.). Ignored when an input LFE is present. */
//         ma_share_mode shareMode;
//     } capture;

//     struct
//     {
//         ma_wasapi_usage usage;              /* When configured, uses Avrt APIs to set the thread characteristics. */
//         ma_bool8 noAutoConvertSRC;          /* When set to true, disables the use of AUDCLNT_STREAMFLAGS_AUTOCONVERTPCM. */
//         ma_bool8 noDefaultQualitySRC;       /* When set to true, disables the use of AUDCLNT_STREAMFLAGS_SRC_DEFAULT_QUALITY. */
//         ma_bool8 noAutoStreamRouting;       /* Disables automatic stream routing. */
//         ma_bool8 noHardwareOffloading;      /* Disables WASAPI's hardware offloading feature. */
//         ma_uint32 loopbackProcessID;        /* The process ID to include or exclude for loopback mode. Set to 0 to capture audio from all processes. Ignored when an explicit device ID is specified. */
//         ma_bool8 loopbackProcessExclude;    /* When set to true, excludes the process specified by loopbackProcessID. By default, the process will be included. */
//     } wasapi;
//     struct
//     {
//         ma_bool32 noMMap;           /* Disables MMap mode. */
//         ma_bool32 noAutoFormat;     /* Opens the ALSA device with SND_PCM_NO_AUTO_FORMAT. */
//         ma_bool32 noAutoChannels;   /* Opens the ALSA device with SND_PCM_NO_AUTO_CHANNELS. */
//         ma_bool32 noAutoResample;   /* Opens the ALSA device with SND_PCM_NO_AUTO_RESAMPLE. */
//     } alsa;
//     struct
//     {
//         const char* pStreamNamePlayback;
//         const char* pStreamNameCapture;
//     } pulse;
//     struct
//     {
//         ma_bool32 allowNominalSampleRateChange; /* Desktop only. When enabled, allows changing of the sample rate at the operating system level. */
//     } coreaudio;
//     struct
//     {
//         ma_opensl_stream_type streamType;
//         ma_opensl_recording_preset recordingPreset;
//         ma_bool32 enableCompatibilityWorkarounds;
//     } opensl;
//     struct
//     {
//         ma_aaudio_usage usage;
//         ma_aaudio_content_type contentType;
//         ma_aaudio_input_preset inputPreset;
//         ma_aaudio_allowed_capture_policy allowedCapturePolicy;
//         ma_bool32 noAutoStartAfterReroute;
//         ma_bool32 enableCompatibilityWorkarounds;
//     } aaudio;
// };
