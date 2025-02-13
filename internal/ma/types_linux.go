//go:build linux

package ma

import "unsafe"

// ma_event
type Event struct {
	Value uint32
	Lock  Mutex
	Cond  unsafe.Pointer
}
