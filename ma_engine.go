package miniaudio

import (
	"unsafe"
)

type (
	maNodeGraph           struct{}
	maResourceManager     struct{}
	maDevice              struct{}
	maLog                 struct{}
	maSpatializerListener struct{}
	maAllocationCallbacks struct{}
	maSoundInlined        struct{}
	maEngineProcessProc   func(userData unsafe.Pointer)
)

type maSpinlock struct {
	// Assume a 4-byte alignment spinlock primitive.
	lock int32
}

type maAtomicUint32 struct {
	value uint32
}

type maEngine struct {
	NodeGraph                          maNodeGraph
	PResourceManager                   *maResourceManager
	PDevice                            *maDevice
	PLog                               *maLog
	SampleRate                         uint32
	ListenerCount                      uint32
	Listeners                          [16]maSpatializerListener // Assuming MA_ENGINE_MAX_LISTENERS = 16
	AllocationCallbacks                maAllocationCallbacks
	OwnsResourceManager                bool
	OwnsDevice                         bool
	InlinedSoundLock                   maSpinlock
	PInlinedSoundHead                  *maSoundInlined
	InlinedSoundCount                  maAtomicUint32
	GainSmoothTimeInFrames             uint32
	DefaultVolumeSmoothTimeInPCMFrames uint32
	MonoExpansionMode                  maMonoExpansionMode
	OnProcess                          maEngineProcessProc
	PProcessUserData                   unsafe.Pointer
}
