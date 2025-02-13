package ma

import "unsafe"

// ma_resampler_config
type ResamplerConfig struct {
	Format           Format                   // ma_format
	Channels         uint32                   // ma_uint32
	SampleRateIn     uint32                   // ma_uint32
	SampleRateOut    uint32                   // ma_uint32
	Algorithm        ResampleAlgorithm        // ma_resample_algorithm
	PBackendVTable   *ResamplingBackendVTable // ma_resampling_backend_vtable*
	PBackendUserData unsafe.Pointer           // void*
	Linear           struct {
		LPFOrder uint32 // ma_uint32
	}
}
