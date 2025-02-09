package miniaudio

import "unsafe"

type ContextConfig struct {
	Log                 *Log                // ma_log*
	ThreadPriority      ThreadPriority      // ma_thread_priority
	ThreadStackSize     uintptr             // size_t
	UserData            unsafe.Pointer      // void*
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
		NoAudioSessionActivate   uint32             // ma_bool32
		NoAudioSessionDeactivate uint32             // ma_bool32
	}

	Jack struct {
		ClientName     *byte  // const char*
		TryStartServer uint32 // ma_bool32
	}

	Custom BackendCallbacks // ma_backend_callbacks
}
