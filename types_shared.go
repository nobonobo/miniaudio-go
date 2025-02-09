package miniaudio

import "unsafe"

type (
	Bool32  uint32         // ma_bool32
	Handle  unsafe.Pointer // ma_handle
	Long    int32          // long
	Mutex   unsafe.Pointer // ma_mutex
	Proc    uintptr        // _proc
	Size    uintptr        // size_t
	Thread  unsafe.Pointer // ma_threat
	VoidPtr unsafe.Pointer // void*
)
