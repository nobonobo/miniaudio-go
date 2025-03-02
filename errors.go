package miniaudio

import "errors"

var (
	ErrDeviceTypeUnknown      = errors.New("miniaudio: device type unknown")
	ErrLibraryNotFound        = errors.New("miniaudio: library not found")
	ErrNotInitialized         = errors.New("miniaudio: library not initialized")
	ErrContextNotInitialized  = errors.New("miniaudio: context not initialized")
	ErrBackendNotSupported    = errors.New("miniaudio: backend not supported")
	ErrSampleRateInvalid      = errors.New("miniaudio: invalid sample rate")
	ErrChannelsInvalid        = errors.New("miniaudio: invalid number of channels")
	ErrNilCallback            = errors.New("miniaudio: nil callback")
	ErrFormatNotSupported     = errors.New("miniaudio: format not supported")
	ErrDeviceNotInitialized   = errors.New("miniaudio: device not initialized")
	ErrDeviceTypeNotSupported = errors.New("miniaudio: device type not supported")
)
