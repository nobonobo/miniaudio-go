package miniaudio

import "unsafe"

type Log struct {
	Callbacks           [MaxLogCallbacks]LogCallback
	CallbackCount       uint32
	AllocationCallbacks AllocationCallbacks
}

type LogCallback struct {
	OnLog    LogCallbackProc
	UserData unsafe.Pointer
}

type LogCallbackProc func(userData unsafe.Pointer, level uint32, message *byte)

type AllocationCallbacks struct {
	UserData  unsafe.Pointer
	OnMalloc  func(size uint, userData unsafe.Pointer) unsafe.Pointer
	OnRealloc func(ptr unsafe.Pointer, size uint, userData unsafe.Pointer) unsafe.Pointer
	OnFree    func(ptr unsafe.Pointer, userData unsafe.Pointer)
}
