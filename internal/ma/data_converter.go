package ma

import "unsafe"

type DataConverter struct {
	FormatIn                Format                     // ma_format
	FormatOut               Format                     // ma_format
	ChannelsIn              uint32                     // ma_uint32
	ChannelsOut             uint32                     // ma_uint32
	SampleRateIn            uint32                     // ma_uint32
	SampleRateOut           uint32                     // ma_uint32
	DitherMode              DitherMode                 // ma_dither_mode
	ExecutionPath           DataConverterExecutionPath // ma_data_converter_execution_path
	ChannelConverter        ChannelConverter           // ma_channel_converter
	Resampler               Resampler                  // ma_resampler
	HasPreFormatConversion  bool                       // ma_bool8
	HasPostFormatConversion bool                       // ma_bool8
	HasChannelConverter     bool                       // ma_bool8
	HasResampler            bool                       // ma_bool8
	IsPassthrough           bool                       // ma_bool8

	// Memory management
	OwnsHeap bool           // ma_bool8
	Heap     unsafe.Pointer // void*
}

// typedef struct
// {
//     ma_format formatIn;
//     ma_format formatOut;
//     ma_uint32 channelsIn;
//     ma_uint32 channelsOut;
//     ma_uint32 sampleRateIn;
//     ma_uint32 sampleRateOut;
//     ma_dither_mode ditherMode;
//     ma_data_converter_execution_path executionPath; /* The execution path the data converter will follow when processing. */
//     ma_channel_converter channelConverter;
//     ma_resampler resampler;
//     ma_bool8 hasPreFormatConversion;
//     ma_bool8 hasPostFormatConversion;
//     ma_bool8 hasChannelConverter;
//     ma_bool8 hasResampler;
//     ma_bool8 isPassthrough;

//     /* Memory management. */
//     ma_bool8 _ownsHeap;
//     void* _pHeap;
// } ma_data_converter;
