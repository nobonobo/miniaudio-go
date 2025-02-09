package miniaudio

// ma_context_config
type ContextConfig struct {
	Log                 *Log                // ma_log*
	ThreadPriority      ThreadPriority      // ma_thread_priority
	ThreadStackSize     Size                // size_t
	UserData            VoidPtr             // void*
	AllocationCallbacks AllocationCallbacks // ma_allocation_callbacks

	ALSA struct {
		UseVerboseDeviceEnumeration Bool32 // ma_bool32
	}

	Pulse struct {
		ApplicationName *byte  // const char*
		ServerName      *byte  // const char*
		TryAutoSpawn    Bool32 // ma_bool32
	}

	CoreAudio struct {
		SessionCategory          IOSSessionCategory // ma_ios_session_category
		SessionCategoryOptions   uint32             // ma_uint32
		NoAudioSessionActivate   Bool32             // ma_bool32
		NoAudioSessionDeactivate Bool32             // ma_bool32
	}

	Jack struct {
		ClientName     *byte  // const char*
		TryStartServer Bool32 // ma_bool32
	}

	Custom BackendCallbacks // ma_backend_callbacks
}
