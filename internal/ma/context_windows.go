//go:build windows

package ma

import (
	"unsafe"
)

// ma_context
type Context struct {
	Callbacks               BackendCallbacks    // ma_backend_callbacks
	Backend                 Backend             // ma_backend
	Log                     *Log                // ma_log*
	LogDetails              Log                 // ma_log
	ThreadPriority          ThreadPriority      // ma_thread_priority
	ThreadStackSize         Size                // size_t
	UserData                unsafe.Pointer      // void*
	AllocationCallbacks     AllocationCallbacks // ma_allocation_callback
	DeviceEnumLock          Mutex               // ma_mutex
	DeviceInfoLock          Mutex               // ma_mutex
	DeviceInfoCapacity      uint32              // ma_uint32
	PlaybackDeviceInfoCount uint32              // ma_uint32
	CaptureDeviceInfoCount  uint32              // ma_uint32
	DeviceInfo              *DeviceInfo         // ma_device_info*

	WASAPI struct {
		CommandThread                   Thread                  // ma_thread (C thread type)
		CommandLock                     Mutex                   // ma_mutex (C mutex)
		CommandSem                      Semaphore               // ma_semaphore (C type)
		CommandIndex                    uint32                  // ma_uint32 (C uint32)
		CommandCount                    uint32                  // ma_uint32 (C uint32)
		Commands                        [4]ContextCommandWASAPI // ma_context_command__wasapi[4] (C array)
		HAvrt                           Handle                  // ma_handle (C handle type)
		AvSetMmThreadCharacteristicsA   Proc                    // ma_proc (C function pointer type)
		AvRevertMmThreadCharacteristics Proc                    // ma_proc (C function pointer type)
		HMMDevapi                       Handle                  // ma_handle (C handle type)
		ActivateAudioInterfaceAsync     Proc                    // ma_proc (C function pointer type)
	}

	Win32 struct {
		HOle32DLL        Handle // ma_handle (C handle type)
		CoInitialize     Proc   // ma_proc (C function pointer type)
		CoInitializeEx   Proc   // ma_proc (C function pointer type)
		CoUninitialize   Proc   // ma_proc (C function pointer type)
		CoCreateInstance Proc   // ma_proc (C function pointer type)
		CoTaskMemFree    Proc   // ma_proc (C function pointer type)
		PropVariantClear Proc   // ma_proc (C function pointer type)
		StringFromGUID2  Proc   // ma_proc (C function pointer type)

		HUser32DLL          Handle // ma_handle (C handle type)
		GetForegroundWindow Proc   // ma_proc (C function pointer type)
		GetDesktopWindow    Proc   // ma_proc (C function pointer type)

		HAdvapi32DLL     Handle // ma_handle (C handle type)
		RegOpenKeyExA    Proc   // ma_proc (C function pointer type)
		RegCloseKey      Proc   // ma_proc (C function pointer type)
		RegQueryValueExA Proc   // ma_proc (C function pointer type)

		CoInitializeResult Long // long (C type)
	}
}

type ContextCommandWASAPI struct {
	Code   int    // int
	PEvent *Event // ma_event*
	Data   struct {
		Quit struct {
			_unused int // int
		}
		CreateAudioClient struct {
			DeviceType         DeviceType      // ma_device_type
			AudioClient        unsafe.Pointer  // void*
			AudioClientService *unsafe.Pointer // void**
			Result             *Result         // ma_result*
		}
		ReleaseAudioClient struct {
			Device     *Device    // ma_device*
			DeviceType DeviceType // ma_device_type
		}
	}
}

// struct ma_context
// {
//     ma_backend_callbacks callbacks;
//     ma_backend backend;                 /* DirectSound, ALSA, etc. */
//     ma_log* pLog;
//     ma_log log; /* Only used if the log is owned by the context. The pLog member will be set to &log in this case. */
//     ma_thread_priority threadPriority;
//     size_t threadStackSize;
//     void* pUserData;
//     ma_allocation_callbacks allocationCallbacks;
//     ma_mutex deviceEnumLock;            /* Used to make ma_context_get_devices() thread safe. */
//     ma_mutex deviceInfoLock;            /* Used to make ma_context_get_device_info() thread safe. */
//     ma_uint32 deviceInfoCapacity;       /* Total capacity of pDeviceInfos. */
//     ma_uint32 playbackDeviceInfoCount;
//     ma_uint32 captureDeviceInfoCount;
//     ma_device_info* pDeviceInfos;       /* Playback devices first, then capture. */

// 	struct
// 	{
// 		ma_thread commandThread;
// 		ma_mutex commandLock;
// 		ma_semaphore commandSem;
// 		ma_uint32 commandIndex;
// 		ma_uint32 commandCount;
// 		ma_context_command__wasapi commands[4];
// 		ma_handle hAvrt;
// 		ma_proc AvSetMmThreadCharacteristicsA;
// 		ma_proc AvRevertMmThreadcharacteristics;
// 		ma_handle hMMDevapi;
// 		ma_proc ActivateAudioInterfaceAsync;
// 	} wasapi;

// 	struct
// 	{
// 		/*HMODULE*/ ma_handle hOle32DLL;
// 		ma_proc CoInitialize;
// 		ma_proc CoInitializeEx;
// 		ma_proc CoUninitialize;
// 		ma_proc CoCreateInstance;
// 		ma_proc CoTaskMemFree;
// 		ma_proc PropVariantClear;
// 		ma_proc StringFromGUID2;

// 		/*HMODULE*/ ma_handle hUser32DLL;
// 		ma_proc GetForegroundWindow;
// 		ma_proc GetDesktopWindow;

// 		/*HMODULE*/ ma_handle hAdvapi32DLL;
// 		ma_proc RegOpenKeyExA;
// 		ma_proc RegCloseKey;
// 		ma_proc RegQueryValueExA;

// 		/*HRESULT*/ long CoInitializeResult;
// 	} win32;
// };
