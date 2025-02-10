package miniaudio

type DataSourceBase struct {
	Vtable         *DataSourceVTable // const ma_data_source_vtable*
	RangeBegFrames uint64            // ma_uint64
	RangeEndFrames uint64            // ma_uint64
	LoopBegFrames  uint64            // ma_uint64
	LoopEndFrames  uint64            // ma_uint64
	Current        DataSource        // ma_data_source*
	Next           DataSource        // ma_data_source*
	OnGetNext      Proc              // ma_data_source_get_next_proc
	IsLooping      Bool32            // MA_ATOMIC(4, ma_bool32)
}

type DataSourceVTable struct {
	OnRead          func(dataSource DataSource, framesOut VoidPtr, frameCount uint64, framesRead *uint64) Result
	OnSeek          func(dataSource DataSource, frameIndex uint64) Result
	OnGetDataFormat func(dataSource DataSource, format *Format, channels *uint32, sampleRate *uint32, channelMap *Channel, channelMapCap Size) Result
	OnGetCursor     func(dataSource DataSource, cursor *uint64) Result
	OnGetLength     func(dataSource DataSource, length *uint64) Result
	OnSetLooping    func(dataSource DataSource, isLooping Bool32) Result
	Flags           uint32
}

// typedef struct
// {
//     const ma_data_source_vtable* vtable;
//     ma_uint64 rangeBegInFrames;
//     ma_uint64 rangeEndInFrames;             /* Set to -1 for unranged (default). */
//     ma_uint64 loopBegInFrames;              /* Relative to rangeBegInFrames. */
//     ma_uint64 loopEndInFrames;              /* Relative to rangeBegInFrames. Set to -1 for the end of the range. */
//     ma_data_source* pCurrent;               /* When non-NULL, the data source being initialized will act as a proxy and will route all operations to pCurrent. Used in conjunction with pNext/onGetNext for seamless chaining. */
//     ma_data_source* pNext;                  /* When set to NULL, onGetNext will be used. */
//     ma_data_source_get_next_proc onGetNext; /* Will be used when pNext is NULL. If both are NULL, no next will be used. */
//     MA_ATOMIC(4, ma_bool32) isLooping;
// } ma_data_source_base;
