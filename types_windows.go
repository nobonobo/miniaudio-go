//go:build windows

package miniaudio

import "unsafe"

// ma_event
type Event unsafe.Pointer

// ma_semaphore
type Semaphore Handle
