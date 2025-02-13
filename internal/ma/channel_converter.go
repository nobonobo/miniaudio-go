package ma

// ma_channel_converter
type ChannelConverter struct {
	Format         Format                // ma_format
	ChannelsIn     uint32                // ma_uint32
	ChannelsOut    uint32                // ma_uint32
	MixingMode     ChannelMixMode        // ma_channel_mix_mode
	ConversionPath ChannelConversionPath // ma_channel_conversion_path
	ChannelMapIn   *Channel              // ma_channel*
	ChannelMapOut  *Channel              // ma_channel*
	ShuffleTable   *uint8                // ma_uint8*
	WeightsF32     **float32             // float**
	WeightsS16     **int32               // ma_int32**
	HeapPointer    VoidPtr               // void*
	OwnsHeap       Bool32                // ma_bool32
}
