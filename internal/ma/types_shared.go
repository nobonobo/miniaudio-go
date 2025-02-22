package ma

import "unsafe"

type (
	Bool32     uint32         // ma_bool32
	DataSource unsafe.Pointer // ma_data_source*
	Handle     unsafe.Pointer // ma_handle
	Long       int32          // long
	Mutex      unsafe.Pointer // ma_mutex
	Proc       uintptr        // _proc
	Size       uintptr        // size_t
	Thread     unsafe.Pointer // ma_threat
)
