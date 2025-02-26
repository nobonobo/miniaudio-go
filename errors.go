package miniaudio

import "errors"

var (
	ErrDeviceTypeUnknown     = errors.New("miniaudio: device type unknown")
	ErrNilPlaybackCallback   = errors.New("miniaudio: nil playback callback")
	ErrNilCaptureCallback    = errors.New("miniaudio: nil capture callback")
	ErrLibraryNotFound       = errors.New("miniaudio: library not found")
	ErrNotInitialized        = errors.New("miniaudio: library not initialized")
	ErrContextNotInitialized = errors.New("miniaudio: context not initialized")
	ErrDeviceNotInitialized  = errors.New("miniaudio: device not initialized")
)
