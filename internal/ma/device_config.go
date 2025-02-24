package ma

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
	DataCallback              Proc               // ma_device_data_proc (function pointer)
	NotificationCallback      Proc               // ma_device_notification_proc (function pointer)
	StopCallback              Proc               // ma_stop_proc (function pointer)
	UserData                  unsafe.Pointer     // void*
	Resampling                ResamplerConfig    // ma_resampler_config

	Playback struct {
		DeviceID                        *DeviceID
		Format                          Format         // ma_format
		Channels                        uint32         // ma_uint32
		ChannelMap                      *Channel       // ma_channel*
		ChannelMixMode                  ChannelMixMode // ma_channel_mix_mode
		CalculateLFEFromSpatialChannels bool           // ma_bool32
		ShareMode                       ShareMode      // ma_share_mode
	}

	Capture struct {
		DeviceID                        *DeviceID
		Format                          Format         // ma_format
		Channels                        uint32         // ma_uint32
		ChannelMap                      *Channel       // *ma_channel
		ChannelMixMode                  ChannelMixMode // ma_channel_mix_mode
		CalculateLFEFromSpatialChannels bool           // ma_bool32
		ShareMode                       ShareMode      // ma_share_mode
	}

	WASAPI struct {
		Usage                  WASAPIUsage // ma_wasapi_usage
		NoAutoConvertSRC       bool        // ma_bool8
		NoDefaultQualitySRC    bool        // ma_bool8
		NoAutoStreamRouting    bool        // ma_bool8
		NoHardwareOffloading   bool        // ma_bool8
		LoopbackProcessID      uint32      // ma_uint32
		LoopbackProcessExclude bool        // ma_bool8
	}

	ALSA struct {
		NoMMap         bool // ma_bool32
		NoAutoFormat   bool // ma_bool32
		NoAutoChannels bool // ma_bool32
		NoAutoResample bool // ma_bool32
	}

	Pulse struct {
		StreamNamePlayback *byte // const char*
		StreamNameCapture  *byte // const char*
	}

	CoreAudio struct {
		AllowNominalSampleRateChange bool // ma_bool32
	}

	OpenSL struct {
		StreamType                     int32 // ma_opensl_stream_type (not implemented)
		RecordingPreset                int32 // ma_opensl_recording_preset (not implemented)
		EnableCompatibilityWorkarounds bool  // ma_bool32
	}

	AAudio struct {
		Usage                          int32 // ma_aaudio_usage (not implemented)
		ContentType                    int32 // ma_aaudio_content_type (not implemented)
		InputPreset                    int32 // ma_aaudio_input_preset (not implemented)
		AllowedCapturePolicy           int32 // ma_aaudio_allowed_capture_policy (not implemented)
		NoAutoStartAfterReroute        bool  // ma_bool32
		EnableCompatibilityWorkarounds bool  // ma_bool32
	}
}
