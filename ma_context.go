package miniaudio

import "unsafe"

type (
	maBackendCallbacks     struct{}
	maBackend              int32
	maThreadPriority       int32
	maMutex                struct{}
	maDeviceInfo           struct{}
	maHandle               unsafe.Pointer
	maProc                 unsafe.Pointer
	maContextCommandWASAPI struct{}
	maThread               struct{}
	maSemaphore            struct{}
)

// ma_context struct translation
type maContext struct {
	Callbacks               maBackendCallbacks
	Backend                 maBackend
	PLog                    *maLog
	Log                     maLog
	ThreadPriority          maThreadPriority
	ThreadStackSize         uintptr
	PUserData               unsafe.Pointer
	AllocationCallbacks     maAllocationCallbacks
	DeviceEnumLock          maMutex
	DeviceInfoLock          maMutex
	DeviceInfoCapacity      uint32
	PlaybackDeviceInfoCount uint32
	CaptureDeviceInfoCount  uint32
	PDeviceInfos            *maDeviceInfo

	// Platform-specific backend fields
	BackendData  interface{}
	PlatformData interface{}
}

// WASAPI backend-specific data
type maContextWASAPI struct {
	CommandThread                   maThread
	CommandLock                     maMutex
	CommandSem                      maSemaphore
	CommandIndex                    uint32
	CommandCount                    uint32
	Commands                        [4]maContextCommandWASAPI
	HAvrt                           maHandle
	AvSetMmThreadCharacteristicsA   maProc
	AvRevertMmThreadCharacteristics maProc
	HMMDevapi                       maHandle
	ActivateAudioInterfaceAsync     maProc
}

// DSOUND backend-specific data
type maContextDSound struct {
	HDSoundDLL                   maHandle
	DirectSoundCreate            maProc
	DirectSoundEnumerateA        maProc
	DirectSoundCaptureCreate     maProc
	DirectSoundCaptureEnumerateA maProc
}

// ALSA backend-specific data
type maContextALSA struct {
	AsoundSO                    maHandle
	SndPcmOpen                  maProc
	SndPcmClose                 maProc
	InternalDeviceEnumLock      maMutex
	UseVerboseDeviceEnumeration bool
}

// PulseAudio backend-specific data
type maContextPulse struct {
	PulseSO          maHandle
	PMainLoop        unsafe.Pointer
	PPulseContext    unsafe.Pointer
	PApplicationName *byte
	PServerName      *byte
}

// JACK backend-specific data
type maContextJACK struct {
	JackSO         maHandle
	PClientName    *byte
	TryStartServer bool
}

// CoreAudio backend-specific data
type maContextCoreAudio struct {
	HCoreFoundation          maHandle
	Component                unsafe.Pointer
	NoAudioSessionDeactivate bool
}

// Platform-specific data for Win32
type maContextWin32 struct {
	HOle32DLL        maHandle
	CoInitialize     maProc
	CoUninitialize   maProc
	CoCreateInstance maProc
}

// Platform-specific data for POSIX
type maContextPOSIX struct {
	Unused int
}
