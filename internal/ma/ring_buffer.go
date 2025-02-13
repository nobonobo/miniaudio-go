package ma

type DuplexRB struct {
	RB PCMRB
}

type PCMRB struct {
	DS         DataSourceBase // ma_data_source_base
	RB         RB             // ma_rb
	Format     Format         // ma_format
	Channels   uint32         // ma_uint32
	SampleRate uint32         // ma_uint32
}

type RB struct {
	Buffer                 VoidPtr             // void*
	SubbufferSizeInBytes   uint32              // ma_uint32
	SubbufferCount         uint32              // ma_uint32
	SubbufferStrideInBytes uint32              // ma_uint32
	EncodedReadOffset      uint32              // MA_ATOMIC(4, ma_uint32)
	EncodedWriteOffset     uint32              // MA_ATOMIC(4, ma_uint32)
	OwnsBuffer             bool                // ma_bool8
	ClearOnWriteAcquire    bool                // ma_bool8
	AllocationCallbacks    AllocationCallbacks // ma_allocation_callbacks
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
