package miniaudio

import (
	"unsafe"
)

type (
	maDeviceID               struct{}
	maVFS                    struct{}
	maDeviceDataProc         func(pDevice unsafe.Pointer, pOutput unsafe.Pointer, frameCount uint32)
	maDeviceNotificationProc func(pDevice unsafe.Pointer, eventType int32)
)

type maEngineConfig struct {
	PResourceManager                   *maResourceManager // Optional resource manager (if defined).
	PContext                           *maContext
	PDevice                            *maDevice
	PPlaybackDeviceID                  *maDeviceID
	DataCallback                       maDeviceDataProc
	NotificationCallback               maDeviceNotificationProc
	PLog                               *maLog
	ListenerCount                      uint32
	Channels                           uint32
	SampleRate                         uint32
	PeriodSizeInFrames                 uint32
	PeriodSizeInMilliseconds           uint32
	GainSmoothTimeInFrames             uint32
	GainSmoothTimeInMilliseconds       uint32
	DefaultVolumeSmoothTimeInPCMFrames uint32
	AllocationCallbacks                maAllocationCallbacks
	NoAutoStart                        bool
	NoDevice                           bool
	MonoExpansionMode                  maMonoExpansionMode
	PResourceManagerVFS                *maVFS
	OnProcess                          maEngineProcessProc
	PProcessUserData                   unsafe.Pointer
}
