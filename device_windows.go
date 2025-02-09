//go:build windows

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

	WASAPI struct {
		AudioClientPlayback              Ptr                    // ma_ptr (C pointer type)
		AudioClientCapture               Ptr                    // ma_ptr (C pointer type)
		RenderClient                     Ptr                    // ma_ptr (C pointer type)
		CaptureClient                    Ptr                    // ma_ptr (C pointer type)
		DeviceEnumerator                 Ptr                    // ma_ptr (C pointer type)
		NotificationClient               IMMNotificationClientr // ma_IMMNotificationClient (C type)
		EventPlayback                    Handle                 // ma_handle (C type)
		EventCapture                     Handle                 // ma_handle (C type)
		ActualBufferSizeInFramesPlayback uint32                 // ma_uint32 (C uint32)
		ActualBufferSizeInFramesCapture  uint32                 // ma_uint32 (C uint32)
		OriginalPeriodSizeInFrames       uint32                 // ma_uint32 (C uint32)
		OriginalPeriodSizeInMilliseconds uint32                 // ma_uint32 (C uint32)
		OriginalPeriods                  uint32                 // ma_uint32 (C uint32)
		OriginalPerformanceProfile       PerformanceProfile     // ma_performance_profile (C type)
		PeriodSizeInFramesPlayback       uint32                 // ma_uint32 (C uint32)
		PeriodSizeInFramesCapture        uint32                 // ma_uint32 (C uint32)
		MappedBufferCapture              unsafe.Pointer         // void* (C pointer type)
		MappedBufferCaptureCap           uint32                 // ma_uint32 (C uint32)
		MappedBufferCaptureLen           uint32                 // ma_uint32 (C uint32)
		MappedBufferPlayback             unsafe.Pointer         // void* (C pointer type)
		MappedBufferPlaybackCap          uint32                 // ma_uint32 (C uint32)
		MappedBufferPlaybackLen          uint32                 // ma_uint32 (C uint32)
		IsStartedCapture                 Bool32                 // ma_atomic_bool32 (C atomic type)
		IsStartedPlayback                Bool32                 // ma_atomic_bool32 (C atomic type)
		LoopbackProcessID                uint32                 // ma_uint32 (C uint32)
		LoopbackProcessExclude           bool                   // ma_bool8 (C bool)
		NoAutoConvertSRC                 bool                   // ma_bool8 (C bool)
		NoDefaultQualitySRC              bool                   // ma_bool8 (C bool)
		NoHardwareOffloading             bool                   // ma_bool8 (C bool)
		AllowCaptureAutoStreamRouting    bool                   // ma_bool8 (C bool)
		AllowPlaybackAutoStreamRouting   bool                   // ma_bool8 (C bool)
		IsDetachedPlayback               bool                   // ma_bool8 (C bool)
		IsDetachedCapture                bool                   // ma_bool8 (C bool)
		Usage                            WASAPIUsage            // ma_wasapi_usage (C enum)
		AvrtHandle                       Handle                 // ma_handle (C type)
		RerouteLock                      Mutex                  // ma_mutex (C mutex type)
	}
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

// 	struct
// 	{
// 		/*IAudioClient**/ ma_ptr pAudioClientPlayback;
// 		/*IAudioClient**/ ma_ptr pAudioClientCapture;
// 		/*IAudioRenderClient**/ ma_ptr pRenderClient;
// 		/*IAudioCaptureClient**/ ma_ptr pCaptureClient;
// 		/*IMMDeviceEnumerator**/ ma_ptr pDeviceEnumerator;      /* Used for IMMNotificationClient notifications. Required for detecting default device changes. */
// 		ma_IMMNotificationClient notificationClient;
// 		/*HANDLE*/ ma_handle hEventPlayback;                    /* Auto reset. Initialized to signaled. */
// 		/*HANDLE*/ ma_handle hEventCapture;                     /* Auto reset. Initialized to unsignaled. */
// 		ma_uint32 actualBufferSizeInFramesPlayback;             /* Value from GetBufferSize(). internalPeriodSizeInFrames is not set to the _actual_ buffer size when low-latency shared mode is being used due to the way the IAudioClient3 API works. */
// 		ma_uint32 actualBufferSizeInFramesCapture;
// 		ma_uint32 originalPeriodSizeInFrames;
// 		ma_uint32 originalPeriodSizeInMilliseconds;
// 		ma_uint32 originalPeriods;
// 		ma_performance_profile originalPerformanceProfile;
// 		ma_uint32 periodSizeInFramesPlayback;
// 		ma_uint32 periodSizeInFramesCapture;
// 		void* pMappedBufferCapture;
// 		ma_uint32 mappedBufferCaptureCap;
// 		ma_uint32 mappedBufferCaptureLen;
// 		void* pMappedBufferPlayback;
// 		ma_uint32 mappedBufferPlaybackCap;
// 		ma_uint32 mappedBufferPlaybackLen;
// 		ma_atomic_bool32 isStartedCapture;                      /* Can be read and written simultaneously across different threads. Must be used atomically, and must be 32-bit. */
// 		ma_atomic_bool32 isStartedPlayback;                     /* Can be read and written simultaneously across different threads. Must be used atomically, and must be 32-bit. */
// 		ma_uint32 loopbackProcessID;
// 		ma_bool8 loopbackProcessExclude;
// 		ma_bool8 noAutoConvertSRC;                              /* When set to true, disables the use of AUDCLNT_STREAMFLAGS_AUTOCONVERTPCM. */
// 		ma_bool8 noDefaultQualitySRC;                           /* When set to true, disables the use of AUDCLNT_STREAMFLAGS_SRC_DEFAULT_QUALITY. */
// 		ma_bool8 noHardwareOffloading;
// 		ma_bool8 allowCaptureAutoStreamRouting;
// 		ma_bool8 allowPlaybackAutoStreamRouting;
// 		ma_bool8 isDetachedPlayback;
// 		ma_bool8 isDetachedCapture;
// 		ma_wasapi_usage usage;
// 		void* hAvrtHandle;
// 		ma_mutex rerouteLock;
// 	} wasapi;
// };
