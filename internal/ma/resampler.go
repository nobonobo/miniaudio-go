package ma

import "unsafe"

// ma_resampler
type Resampler struct {
	Backend         *ResamplingBackend       // ma_resampling_backend*
	BackendVTable   *ResamplingBackendVTable // ma_resampling_backend_vtable*
	BackendUserData unsafe.Pointer           // void*
	Format          Format                   // ma_format
	Channels        uint32                   // ma_uint32
	SampleRateIn    uint32                   // ma_uint32
	SampleRateOut   uint32                   // ma_uint32
	State           struct {
		Linear LinearResampler // ma_linear_resampler
	}
	Heap     unsafe.Pointer // void*
	OwnsHeap Bool32         // ma_bool32
}

// ma_linear_resampler
type LinearResampler struct {
	Config        LinearResamplerConfig // ma_linear_resampler_config
	InAdvanceInt  uint32                // ma_uint32
	InAdvanceFrac uint32                // ma_uint32
	InTimeInt     uint32                // ma_uint32
	InTimeFrac    uint32                // ma_uint32
	X0            struct {
		F32 *float32 // float*
		S16 *int16   // ma_int16*
	} // union x0
	X1 struct {
		F32 *float32 // float*
		S16 *int16   // ma_int16*
	} // union x1
	LPF      LPF            // ma_lpf
	Heap     unsafe.Pointer // void*
	OwnsHeap Bool32         // ma_bool32
}

// ma_linear_resampler_config
type LinearResamplerConfig struct {
	Format           Format  // ma_format
	Channels         uint32  // ma_uint32
	SampleRateIn     uint32  // ma_uint32
	SampleRateOut    uint32  // ma_uint32
	LPFOrder         uint32  // ma_uint32
	LPFNyquistFactor float64 // double
}

// ma_lpf
type LPF struct {
	Format     Format         // ma_format
	Channels   uint32         // ma_uint32
	SampleRate uint32         // ma_uint32
	LPF1Count  uint32         // ma_uint32
	LPF2Count  uint32         // ma_uint32
	LPF1       *LPF1          // *ma_lpf1
	LPF2       *LPF2          // *ma_lpf2
	Heap       unsafe.Pointer // void*
	OwnsHeap   Bool32         // ma_bool32
}

// ma_lpf1
type LPF1 struct {
	Format   Format             // ma_format
	Channels uint32             // ma_uint32
	A        BiquadCoefficient  // ma_biquad_coefficient
	R1       *BiquadCoefficient // *ma_biquad_coefficient
	Heap     unsafe.Pointer     // void*
	OwnsHeap Bool32             // ma_bool32
}

// ma_biquad_coefficient
type BiquadCoefficient struct {
	AsFloat32 float32 // float
	AsInt32   int32   // int32
}

// ma_lpf2
type LPF2 struct {
	BQ Biquad // ma_biquad
}

// ma_biquad
type Biquad struct {
	Format   Format             // ma_format
	Channels uint32             // ma_uint32
	B0       BiquadCoefficient  // ma_biquad_coefficient
	B1       BiquadCoefficient  // ma_biquad_coefficient
	B2       BiquadCoefficient  // ma_biquad_coefficient
	A1       BiquadCoefficient  // ma_biquad_coefficient
	A2       BiquadCoefficient  // ma_biquad_coefficient
	R1       *BiquadCoefficient // ma_biquad_coefficient*
	R2       *BiquadCoefficient // ma_biquad_coefficient*
	Heap     unsafe.Pointer     // void*
	OwnsHeap Bool32             // ma_bool32
}
