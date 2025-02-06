package miniaudio

import "unsafe"

// ma_resampling_backend
type ResamplingBackend struct{}

// ma_resampling_backend_vtable
type ResamplingBackendVTable struct {
	OnGetHeapSize                 func(pUserData unsafe.Pointer, pConfig *ResamplerConfig, pHeapSizeInBytes *uint64) Result
	OnInit                        func(pUserData unsafe.Pointer, pConfig *ResamplerConfig, pHeap unsafe.Pointer, ppBackend **ResamplingBackend) Result
	OnUninit                      func(pUserData unsafe.Pointer, pBackend *ResamplingBackend, pAllocationCallbacks unsafe.Pointer)
	OnProcess                     func(pUserData unsafe.Pointer, pBackend *ResamplingBackend, pFramesIn unsafe.Pointer, pFrameCountIn *uint64, pFramesOut unsafe.Pointer, pFrameCountOut *uint64) Result
	OnSetRate                     func(pUserData unsafe.Pointer, pBackend *ResamplingBackend, sampleRateIn uint32, sampleRateOut uint32) Result
	OnGetInputLatency             func(pUserData unsafe.Pointer, pBackend *ResamplingBackend) uint64
	OnGetOutputLatency            func(pUserData unsafe.Pointer, pBackend *ResamplingBackend) uint64
	OnGetRequiredInputFrameCount  func(pUserData unsafe.Pointer, pBackend *ResamplingBackend, outputFrameCount uint64, pInputFrameCount *uint64) Result
	OnGetExpectedOutputFrameCount func(pUserData unsafe.Pointer, pBackend *ResamplingBackend, inputFrameCount uint64, pOutputFrameCount *uint64) Result
	OnReset                       func(pUserData unsafe.Pointer, pBackend *ResamplingBackend) Result
}
