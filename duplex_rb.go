package miniaudio

type DuplexRB struct {
	RB PCMRB
}

// TODO: continue here!
type PCMRB struct {
	DS         DataSourceBase // ma_data_source_base
	RB         RB             // ma_rb
	Format     Format         // ma_format
	Channels   uint32         // ma_uint32
	SampleRate uint32         // ma_uint32
}

// typedef struct
// {
//     void* pBuffer;
//     ma_uint32 subbufferSizeInBytes;
//     ma_uint32 subbufferCount;
//     ma_uint32 subbufferStrideInBytes;
//     MA_ATOMIC(4, ma_uint32) encodedReadOffset;  /* Most significant bit is the loop flag. Lower 31 bits contains the actual offset in bytes. Must be used atomically. */
//     MA_ATOMIC(4, ma_uint32) encodedWriteOffset; /* Most significant bit is the loop flag. Lower 31 bits contains the actual offset in bytes. Must be used atomically. */
//     ma_bool8 ownsBuffer;                        /* Used to know whether or not miniaudio is responsible for free()-ing the buffer. */
//     ma_bool8 clearOnWriteAcquire;               /* When set, clears the acquired write buffer before returning from ma_rb_acquire_write(). */
//     ma_allocation_callbacks allocationCallbacks;
// } ma_rb;

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
