//go:build linux

package miniaudio

import (
	"unsafe"
)

// ma_device
type Device struct {
	Context                   *Context    // ma_context*
	DeviceType                DeviceType  // ma_device_type
	SampleRate                uint32      // ma_uint32
	State                     DeviceState // ma_atomic_device_state
	OnData                    Proc        // ma_device_data_proc
	OnNotification            Proc        // ma_device_notification_proc
	OnStop                    Proc        // ma_stop_proc
	UserData                  VoidPtr     // void*
	StartStopLock             Mutex       // ma_mutex
	WakeupEvent               Event       // ma_event
	StartEvent                Event       // ma_event
	StopEvent                 Event       // ma_event
	Thread                    Thread      // ma_thread
	WorkResult                Result      // ma_result
	IsOwnerOfContext          bool        // ma_bool8
	NoPreSilencedOutputBuffer bool        // ma_bool8
	NoClip                    bool        // ma_bool8
	NoDisableDenormals        bool        // ma_bool8
	NoFixedSizedCallback      bool        // ma_bool8
	MasterVolumeFactor        float32     // ma_atomic_float
	DuplexRB                  DuplexRB    // ma_duplex_rb

	Resampling struct {
		Algorithm       ResampleAlgorithm        // ma_resample_algorithm
		BackendVTable   *ResamplingBackendVTable // ma_resampling_backend_vtable*
		BackendUserData VoidPtr                  // void*
		Linear          struct {
			LPFOrder uint32 // ma_uint32
		}
	}

	Playback struct {
		ID                              *DeviceID                     // ma_device_id* (C pointer type)
		DeviceID                        DeviceID                      // ma_device_id (C type)
		Name                            [MaxDeviceNameLength + 1]byte // char[MA_MAX_DEVICE_NAME_LENGTH+1] (C array)
		ShareMode                       ShareMode                     // ma_share_mode (C enum)
		Format                          Format                        // ma_format (C enum)
		Channels                        uint32                        // ma_uint32 (C uint32)
		ChannelMap                      [MaxChannels]Channel          // ma_channel[MA_MAX_CHANNELS] (C array)
		InternalFormat                  Format                        // ma_format (C enum)
		InternalChannels                uint32                        // ma_uint32 (C uint32)
		InternalSampleRate              uint32                        // ma_uint32 (C uint32)
		InternalChannelMap              [MaxChannels]Channel          // ma_channel[MA_MAX_CHANNELS] (C array)
		InternalPeriodSizeInFrames      uint32                        // ma_uint32 (C uint32)
		InternalPeriods                 uint32                        // ma_uint32 (C uint32)
		ChannelMixMode                  ChannelMixMode                // ma_channel_mix_mode (C enum)
		CalculateLFEFromSpatialChannels bool                          // ma_bool32 (C bool)
		Converter                       DataConverter                 // ma_data_converter (C type)
		IntermediaryBuffer              unsafe.Pointer                // void* (C pointer type)
		IntermediaryBufferCap           uint32                        // ma_uint32 (C uint32)
		IntermediaryBufferLen           uint32                        // ma_uint32 (C uint32)
		InputCache                      unsafe.Pointer                // void* (C pointer type)
		InputCacheCap                   uint64                        // ma_uint64 (C uint64)
		InputCacheConsumed              uint64                        // ma_uint64 (C uint64)
		InputCacheRemaining             uint64                        // ma_uint64 (C uint64)
	}

	Capture struct {
		ID                              *DeviceID                     // ma_device_id* (C pointer type)
		DeviceID                        DeviceID                      // ma_device_id (C type)
		Name                            [MaxDeviceNameLength + 1]byte // char[MA_MAX_DEVICE_NAME_LENGTH+1] (C array)
		ShareMode                       ShareMode                     // ma_share_mode (C enum)
		Format                          Format                        // ma_format (C enum)
		Channels                        uint32                        // ma_uint32 (C uint32)
		ChannelMap                      [MaxChannels]Channel          // ma_channel[MA_MAX_CHANNELS] (C array)
		InternalFormat                  Format                        // ma_format (C enum)
		InternalChannels                uint32                        // ma_uint32 (C uint32)
		InternalSampleRate              uint32                        // ma_uint32 (C uint32)
		InternalChannelMap              [MaxChannels]Channel          // ma_channel[MA_MAX_CHANNELS] (C array)
		InternalPeriodSizeInFrames      uint32                        // ma_uint32 (C uint32)
		InternalPeriods                 uint32                        // ma_uint32 (C uint32)
		ChannelMixMode                  ChannelMixMode                // ma_channel_mix_mode (C enum)
		CalculateLFEFromSpatialChannels bool                          // ma_bool32 (C bool)
		Converter                       DataConverter                 // ma_data_converter (C type)
		IntermediaryBuffer              unsafe.Pointer                // void* (C pointer type)
		IntermediaryBufferCap           uint32                        // ma_uint32 (C uint32)
		IntermediaryBufferLen           uint32                        // ma_uint32 (C uint32)
	}

	BackendUnion unsafe.Pointer
}

type DeviceALSA struct {
	PCMPlayback                 uintptr        // snd_pcm_t*
	PCMCapture                  uintptr        // snd_pcm_t*
	PollDescriptorsPlayback     unsafe.Pointer // *struct pollfd
	PollDescriptorsCapture      unsafe.Pointer // *struct pollfd
	PollDescriptorCountPlayback int
	PollDescriptorCountCapture  int
	WakeupfdPlayback            int  // eventfd for waking up from poll() when the playback device is stopped
	WakeupfdCapture             int  // eventfd for waking up from poll() when the capture device is stopped
	IsUsingMMapPlayback         bool // ma_bool8
	IsUsingMMapCapture          bool // ma_bool8
}

type DevicePulse struct {
	MainLoop       uintptr // pa_mainloop*
	PulseContext   uintptr // pa_context*
	StreamPlayback uintptr // pa_stream*
	StreamCapture  uintptr // pa_stream*
}

type DeviceJack struct {
	Client                     uintptr  // jack_client_t*
	PortsPlayback              *uintptr // *jack_port_t**
	PortsCapture               *uintptr // *jack_port_t**
	IntermediaryBufferPlayback *float32 // Always float32 since JACK uses floating point
	IntermediaryBufferCapture  *float32 // Always float32
}

// struct ma_device
// {
//     ma_context* pContext;
//     ma_device_type type;
//     ma_uint32 sampleRate;
//     ma_atomic_device_state state;               /* The state of the device is variable and can change at any time on any thread. Must be used atomically. */
//     ma_device_data_proc onData;                 /* Set once at initialization time and should not be changed after. */
//     ma_device_notification_proc onNotification; /* Set once at initialization time and should not be changed after. */
//     ma_stop_proc onStop;                        /* DEPRECATED. Use the notification callback instead. Set once at initialization time and should not be changed after. */
//     void* pUserData;                            /* Application defined data. */
//     ma_mutex startStopLock;
//     ma_event wakeupEvent;
//     ma_event startEvent;
//     ma_event stopEvent;
//     ma_thread thread;
//     ma_result workResult;                       /* This is set by the worker thread after it's finished doing a job. */
//     ma_bool8 isOwnerOfContext;                  /* When set to true, uninitializing the device will also uninitialize the context. Set to true when NULL is passed into ma_device_init(). */
//     ma_bool8 noPreSilencedOutputBuffer;
//     ma_bool8 noClip;
//     ma_bool8 noDisableDenormals;
//     ma_bool8 noFixedSizedCallback;
//     ma_atomic_float masterVolumeFactor;         /* Linear 0..1. Can be read and written simultaneously by different threads. Must be used atomically. */
//     ma_duplex_rb duplexRB;                      /* Intermediary buffer for duplex device on asynchronous backends. */
//     struct
//     {
//         ma_resample_algorithm algorithm;
//         ma_resampling_backend_vtable* pBackendVTable;
//         void* pBackendUserData;
//         struct
//         {
//             ma_uint32 lpfOrder;
//         } linear;
//     } resampling;
//     struct
//     {
//         ma_device_id* pID;                  /* Set to NULL if using default ID, otherwise set to the address of "id". */
//         ma_device_id id;                    /* If using an explicit device, will be set to a copy of the ID used for initialization. Otherwise cleared to 0. */
//         char name[MA_MAX_DEVICE_NAME_LENGTH + 1];                     /* Maybe temporary. Likely to be replaced with a query API. */
//         ma_share_mode shareMode;            /* Set to whatever was passed in when the device was initialized. */
//         ma_format format;
//         ma_uint32 channels;
//         ma_channel channelMap[MA_MAX_CHANNELS];
//         ma_format internalFormat;
//         ma_uint32 internalChannels;
//         ma_uint32 internalSampleRate;
//         ma_channel internalChannelMap[MA_MAX_CHANNELS];
//         ma_uint32 internalPeriodSizeInFrames;
//         ma_uint32 internalPeriods;
//         ma_channel_mix_mode channelMixMode;
//         ma_bool32 calculateLFEFromSpatialChannels;
//         ma_data_converter converter;
//         void* pIntermediaryBuffer;          /* For implementing fixed sized buffer callbacks. Will be null if using variable sized callbacks. */
//         ma_uint32 intermediaryBufferCap;
//         ma_uint32 intermediaryBufferLen;    /* How many valid frames are sitting in the intermediary buffer. */
//         void* pInputCache;                  /* In external format. Can be null. */
//         ma_uint64 inputCacheCap;
//         ma_uint64 inputCacheConsumed;
//         ma_uint64 inputCacheRemaining;
//     } playback;
//     struct
//     {
//         ma_device_id* pID;                  /* Set to NULL if using default ID, otherwise set to the address of "id". */
//         ma_device_id id;                    /* If using an explicit device, will be set to a copy of the ID used for initialization. Otherwise cleared to 0. */
//         char name[MA_MAX_DEVICE_NAME_LENGTH + 1];                     /* Maybe temporary. Likely to be replaced with a query API. */
//         ma_share_mode shareMode;            /* Set to whatever was passed in when the device was initialized. */
//         ma_format format;
//         ma_uint32 channels;
//         ma_channel channelMap[MA_MAX_CHANNELS];
//         ma_format internalFormat;
//         ma_uint32 internalChannels;
//         ma_uint32 internalSampleRate;
//         ma_channel internalChannelMap[MA_MAX_CHANNELS];
//         ma_uint32 internalPeriodSizeInFrames;
//         ma_uint32 internalPeriods;
//         ma_channel_mix_mode channelMixMode;
//         ma_bool32 calculateLFEFromSpatialChannels;
//         ma_data_converter converter;
//         void* pIntermediaryBuffer;          /* For implementing fixed sized buffer callbacks. Will be null if using variable sized callbacks. */
//         ma_uint32 intermediaryBufferCap;
//         ma_uint32 intermediaryBufferLen;    /* How many valid frames are sitting in the intermediary buffer. */
//     } capture;

//     union
//     {
// #ifdef MA_SUPPORT_ALSA
//         struct
//         {
//             /*snd_pcm_t**/ ma_ptr pPCMPlayback;
//             /*snd_pcm_t**/ ma_ptr pPCMCapture;
//             /*struct pollfd**/ void* pPollDescriptorsPlayback;
//             /*struct pollfd**/ void* pPollDescriptorsCapture;
//             int pollDescriptorCountPlayback;
//             int pollDescriptorCountCapture;
//             int wakeupfdPlayback;   /* eventfd for waking up from poll() when the playback device is stopped. */
//             int wakeupfdCapture;    /* eventfd for waking up from poll() when the capture device is stopped. */
//             ma_bool8 isUsingMMapPlayback;
//             ma_bool8 isUsingMMapCapture;
//         } alsa;
// #endif
// #ifdef MA_SUPPORT_PULSEAUDIO
//         struct
//         {
//             /*pa_mainloop**/ ma_ptr pMainLoop;
//             /*pa_context**/ ma_ptr pPulseContext;
//             /*pa_stream**/ ma_ptr pStreamPlayback;
//             /*pa_stream**/ ma_ptr pStreamCapture;
//         } pulse;
// #endif
// #ifdef MA_SUPPORT_JACK
//         struct
//         {
//             /*jack_client_t**/ ma_ptr pClient;
//             /*jack_port_t**/ ma_ptr* ppPortsPlayback;
//             /*jack_port_t**/ ma_ptr* ppPortsCapture;
//             float* pIntermediaryBufferPlayback; /* Typed as a float because JACK is always floating point. */
//             float* pIntermediaryBufferCapture;
//         } jack;
// #endif
//     };
// };
