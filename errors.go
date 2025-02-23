package miniaudio

import "errors"

var (
	ErrLibraryNotFound      = errors.New("miniaudio: library not found")
	ErrNotInitialized       = errors.New("miniaudio: library not initialized")
	ErrDeviceNotInitialized = errors.New("miniaudio: device not initialized")
)
