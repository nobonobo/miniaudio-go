package miniaudio

import (
	"unsafe"

	"github.com/samborkent/miniaudio/internal/ma"
)

type (
	PlaybackCallback[T Formats]  func(frameCount, channelCount int) [][]T
	CaptureCallback[T Formats]   func(inputSamples []T, frameCount, channelCount int)
	DuplexCallback[P, C Formats] func(inputSamples []C, frameCount, playbackChannels, captureChannels int) [][]P
)

func SetPlaybackCallback[T Formats](config *DeviceConfig, callback PlaybackCallback[T]) {
	config.dataCallback.Store(uintptr(unsafe.Pointer(&callback)))
}

func SetCaptureCallback[T Formats](config *DeviceConfig, callback CaptureCallback[T]) {
	config.dataCallback.Store(uintptr(unsafe.Pointer(&callback)))
}

func SetDuplexCallback[P, C Formats](config *DeviceConfig, callback DuplexCallback[P, C]) {
	config.dataCallback.Store(uintptr(unsafe.Pointer(&callback)))
}

type dataCallback func(device *ma.Device, output, input unsafe.Pointer, frameCount uint32) uintptr

func (c *DeviceConfig) playbackCallback() (dataCallback, error) {
	if c.dataCallback.Load() == 0 {
		return nil, ErrNilCallback
	}

	switch c.Playback.Format {
	case FormatUint8:
		callbackPtr := (*PlaybackCallback[uint8])(unsafe.Pointer(c.dataCallback.Load()))
		if callbackPtr == nil {
			return nil, ErrNilCallback
		}

		callback := *callbackPtr

		return func(_ *ma.Device, output, _ unsafe.Pointer, frameCount uint32) uintptr {
			outputSamples := unsafe.Slice((*uint8)(output), frameCount*uint32(c.Playback.Channels))

			gotSamples := callback(int(frameCount), c.Playback.Channels)

			for i := range int(frameCount) {
				for j := range c.Playback.Channels {
					outputSamples[i*c.Playback.Channels+j] = gotSamples[i][j]
				}
			}

			return 0
		}, nil
	case FormatInt16:
		callbackPtr := (*PlaybackCallback[int16])(unsafe.Pointer(c.dataCallback.Load()))
		if callbackPtr == nil {
			return nil, ErrNilCallback
		}

		callback := *callbackPtr

		return func(_ *ma.Device, output, _ unsafe.Pointer, frameCount uint32) uintptr {
			outputSamples := unsafe.Slice((*int16)(output), frameCount*uint32(c.Playback.Channels))

			gotSamples := callback(int(frameCount), c.Playback.Channels)

			for i := range int(frameCount) {
				for j := range c.Playback.Channels {
					outputSamples[i*c.Playback.Channels+j] = gotSamples[i][j]
				}
			}

			return 0
		}, nil
	case FormatInt32:
		callbackPtr := (*PlaybackCallback[int32])(unsafe.Pointer(c.dataCallback.Load()))
		if callbackPtr == nil {
			return nil, ErrNilCallback
		}

		callback := *callbackPtr

		return func(_ *ma.Device, output, _ unsafe.Pointer, frameCount uint32) uintptr {
			outputSamples := unsafe.Slice((*int32)(output), frameCount*uint32(c.Playback.Channels))

			gotSamples := callback(int(frameCount), c.Playback.Channels)

			for i := range int(frameCount) {
				for j := range c.Playback.Channels {
					outputSamples[i*c.Playback.Channels+j] = gotSamples[i][j]
				}
			}

			return 0
		}, nil
	case FormatFloat32:
		callbackPtr := (*PlaybackCallback[float32])(unsafe.Pointer(c.dataCallback.Load()))
		if callbackPtr == nil {
			return nil, ErrNilCallback
		}

		callback := *callbackPtr

		return func(_ *ma.Device, output, _ unsafe.Pointer, frameCount uint32) uintptr {
			outputSamples := unsafe.Slice((*float32)(output), frameCount*uint32(c.Playback.Channels))

			gotSamples := callback(int(frameCount), c.Playback.Channels)

			for i := range int(frameCount) {
				for j := range c.Playback.Channels {
					outputSamples[i*c.Playback.Channels+j] = gotSamples[i][j]
				}
			}

			return 0
		}, nil
	default:
		return nil, ErrFormatNotSupported
	}
}

func (c *DeviceConfig) captureCallback() (dataCallback, error) {
	if c.dataCallback.Load() == 0 {
		return nil, ErrNilCallback
	}

	switch c.Capture.Format {
	case FormatUint8:
		callbackPtr := (*CaptureCallback[uint8])(unsafe.Pointer(c.dataCallback.Load()))
		if callbackPtr == nil {
			return nil, ErrNilCallback
		}

		callback := *callbackPtr

		return func(_ *ma.Device, _, input unsafe.Pointer, frameCount uint32) uintptr {
			inputSamples := unsafe.Slice((*uint8)(input), int(frameCount)*c.Capture.Channels)

			callback(inputSamples, int(frameCount), c.Capture.Channels)

			return 0
		}, nil
	case FormatInt16:
		callbackPtr := (*CaptureCallback[int16])(unsafe.Pointer(c.dataCallback.Load()))
		if callbackPtr == nil {
			return nil, ErrNilCallback
		}

		callback := *callbackPtr

		return func(_ *ma.Device, _, input unsafe.Pointer, frameCount uint32) uintptr {
			inputSamples := unsafe.Slice((*int16)(input), int(frameCount)*c.Capture.Channels)

			callback(inputSamples, int(frameCount), c.Capture.Channels)

			return 0
		}, nil
	case FormatInt32:
		callbackPtr := (*CaptureCallback[int32])(unsafe.Pointer(c.dataCallback.Load()))
		if callbackPtr == nil {
			return nil, ErrNilCallback
		}

		callback := *callbackPtr

		return func(_ *ma.Device, _, input unsafe.Pointer, frameCount uint32) uintptr {
			inputSamples := unsafe.Slice((*int32)(input), int(frameCount)*c.Capture.Channels)

			callback(inputSamples, int(frameCount), c.Capture.Channels)

			return 0
		}, nil
	case FormatFloat32:
		callbackPtr := (*CaptureCallback[float32])(unsafe.Pointer(c.dataCallback.Load()))
		if callbackPtr == nil {
			return nil, ErrNilCallback
		}

		callback := *callbackPtr

		return func(_ *ma.Device, _, input unsafe.Pointer, frameCount uint32) uintptr {
			inputSamples := unsafe.Slice((*float32)(input), int(frameCount)*c.Capture.Channels)

			callback(inputSamples, int(frameCount), c.Capture.Channels)

			return 0
		}, nil
	default:
		return nil, ErrFormatNotSupported
	}
}

func (c *DeviceConfig) duplexCallback() (dataCallback, error) {
	if c.dataCallback.Load() == 0 {
		return nil, ErrNilCallback
	}

	switch c.Playback.Format {
	case FormatUint8:
		switch c.Capture.Format {
		case FormatUint8:
			callbackPtr := (*DuplexCallback[uint8, uint8])(unsafe.Pointer(c.dataCallback.Load()))
			if callbackPtr == nil {
				return nil, ErrNilCallback
			}

			callback := *callbackPtr

			return func(_ *ma.Device, output, input unsafe.Pointer, frameCount uint32) uintptr {
				inputSamples := unsafe.Slice((*uint8)(input), int(frameCount)*c.Capture.Channels)
				outputSamples := unsafe.Slice((*uint8)(output), int(frameCount)*c.Playback.Channels)

				gotSamples := callback(inputSamples, int(frameCount), c.Playback.Channels, c.Capture.Channels)

				for i := range int(frameCount) {
					for j := range c.Playback.Channels {
						outputSamples[i*c.Playback.Channels+j] = gotSamples[i][j]
					}
				}

				return 0
			}, nil
		case FormatInt16:
			callbackPtr := (*DuplexCallback[uint8, int16])(unsafe.Pointer(c.dataCallback.Load()))
			if callbackPtr == nil {
				return nil, ErrNilCallback
			}

			callback := *callbackPtr

			return func(_ *ma.Device, output, input unsafe.Pointer, frameCount uint32) uintptr {
				inputSamples := unsafe.Slice((*int16)(input), int(frameCount)*c.Capture.Channels)
				outputSamples := unsafe.Slice((*uint8)(output), int(frameCount)*c.Playback.Channels)

				gotSamples := callback(inputSamples, int(frameCount), c.Playback.Channels, c.Capture.Channels)

				for i := range int(frameCount) {
					for j := range c.Playback.Channels {
						outputSamples[i*c.Playback.Channels+j] = gotSamples[i][j]
					}
				}

				return 0
			}, nil
		case FormatInt32:
			callbackPtr := (*DuplexCallback[uint8, int32])(unsafe.Pointer(c.dataCallback.Load()))
			if callbackPtr == nil {
				return nil, ErrNilCallback
			}

			callback := *callbackPtr

			return func(_ *ma.Device, output, input unsafe.Pointer, frameCount uint32) uintptr {
				inputSamples := unsafe.Slice((*int32)(input), int(frameCount)*c.Capture.Channels)
				outputSamples := unsafe.Slice((*uint8)(output), int(frameCount)*c.Playback.Channels)

				gotSamples := callback(inputSamples, int(frameCount), c.Playback.Channels, c.Capture.Channels)

				for i := range int(frameCount) {
					for j := range c.Playback.Channels {
						outputSamples[i*c.Playback.Channels+j] = gotSamples[i][j]
					}
				}

				return 0
			}, nil
		case FormatFloat32:
			callbackPtr := (*DuplexCallback[uint8, float32])(unsafe.Pointer(c.dataCallback.Load()))
			if callbackPtr == nil {
				return nil, ErrNilCallback
			}

			callback := *callbackPtr

			return func(_ *ma.Device, output, input unsafe.Pointer, frameCount uint32) uintptr {
				inputSamples := unsafe.Slice((*float32)(input), int(frameCount)*c.Capture.Channels)
				outputSamples := unsafe.Slice((*uint8)(output), int(frameCount)*c.Playback.Channels)

				gotSamples := callback(inputSamples, int(frameCount), c.Playback.Channels, c.Capture.Channels)

				for i := range int(frameCount) {
					for j := range c.Playback.Channels {
						outputSamples[i*c.Playback.Channels+j] = gotSamples[i][j]
					}
				}

				return 0
			}, nil
		default:
			return nil, ErrFormatNotSupported
		}
	case FormatInt16:
		switch c.Capture.Format {
		case FormatUint8:
			callbackPtr := (*DuplexCallback[int16, uint8])(unsafe.Pointer(c.dataCallback.Load()))
			if callbackPtr == nil {
				return nil, ErrNilCallback
			}

			callback := *callbackPtr

			return func(_ *ma.Device, output, input unsafe.Pointer, frameCount uint32) uintptr {
				inputSamples := unsafe.Slice((*uint8)(input), int(frameCount)*c.Capture.Channels)
				outputSamples := unsafe.Slice((*int16)(output), int(frameCount)*c.Playback.Channels)

				gotSamples := callback(inputSamples, int(frameCount), c.Playback.Channels, c.Capture.Channels)

				for i := range int(frameCount) {
					for j := range c.Playback.Channels {
						outputSamples[i*c.Playback.Channels+j] = gotSamples[i][j]
					}
				}

				return 0
			}, nil
		case FormatInt16:
			callbackPtr := (*DuplexCallback[int16, int16])(unsafe.Pointer(c.dataCallback.Load()))
			if callbackPtr == nil {
				return nil, ErrNilCallback
			}

			callback := *callbackPtr

			return func(_ *ma.Device, output, input unsafe.Pointer, frameCount uint32) uintptr {
				inputSamples := unsafe.Slice((*int16)(input), int(frameCount)*c.Capture.Channels)
				outputSamples := unsafe.Slice((*int16)(output), int(frameCount)*c.Playback.Channels)

				gotSamples := callback(inputSamples, int(frameCount), c.Playback.Channels, c.Capture.Channels)

				for i := range int(frameCount) {
					for j := range c.Playback.Channels {
						outputSamples[i*c.Playback.Channels+j] = gotSamples[i][j]
					}
				}

				return 0
			}, nil
		case FormatInt32:
			callbackPtr := (*DuplexCallback[int16, int32])(unsafe.Pointer(c.dataCallback.Load()))
			if callbackPtr == nil {
				return nil, ErrNilCallback
			}

			callback := *callbackPtr

			return func(_ *ma.Device, output, input unsafe.Pointer, frameCount uint32) uintptr {
				inputSamples := unsafe.Slice((*int32)(input), int(frameCount)*c.Capture.Channels)
				outputSamples := unsafe.Slice((*int16)(output), int(frameCount)*c.Playback.Channels)

				gotSamples := callback(inputSamples, int(frameCount), c.Playback.Channels, c.Capture.Channels)

				for i := range int(frameCount) {
					for j := range c.Playback.Channels {
						outputSamples[i*c.Playback.Channels+j] = gotSamples[i][j]
					}
				}

				return 0
			}, nil
		case FormatFloat32:
			callbackPtr := (*DuplexCallback[int16, float32])(unsafe.Pointer(c.dataCallback.Load()))
			if callbackPtr == nil {
				return nil, ErrNilCallback
			}

			callback := *callbackPtr

			return func(_ *ma.Device, output, input unsafe.Pointer, frameCount uint32) uintptr {
				inputSamples := unsafe.Slice((*float32)(input), int(frameCount)*c.Capture.Channels)
				outputSamples := unsafe.Slice((*int16)(output), int(frameCount)*c.Playback.Channels)

				gotSamples := callback(inputSamples, int(frameCount), c.Playback.Channels, c.Capture.Channels)

				for i := range int(frameCount) {
					for j := range c.Playback.Channels {
						outputSamples[i*c.Playback.Channels+j] = gotSamples[i][j]
					}
				}

				return 0
			}, nil
		default:
			return nil, ErrFormatNotSupported
		}
	case FormatInt32:
		switch c.Capture.Format {
		case FormatUint8:
			callbackPtr := (*DuplexCallback[int32, uint8])(unsafe.Pointer(c.dataCallback.Load()))
			if callbackPtr == nil {
				return nil, ErrNilCallback
			}

			callback := *callbackPtr

			return func(_ *ma.Device, output, input unsafe.Pointer, frameCount uint32) uintptr {
				inputSamples := unsafe.Slice((*uint8)(input), int(frameCount)*c.Capture.Channels)
				outputSamples := unsafe.Slice((*int32)(output), int(frameCount)*c.Playback.Channels)

				gotSamples := callback(inputSamples, int(frameCount), c.Playback.Channels, c.Capture.Channels)

				for i := range int(frameCount) {
					for j := range c.Playback.Channels {
						outputSamples[i*c.Playback.Channels+j] = gotSamples[i][j]
					}
				}

				return 0
			}, nil
		case FormatInt16:
			callbackPtr := (*DuplexCallback[int32, int16])(unsafe.Pointer(c.dataCallback.Load()))
			if callbackPtr == nil {
				return nil, ErrNilCallback
			}

			callback := *callbackPtr

			return func(_ *ma.Device, output, input unsafe.Pointer, frameCount uint32) uintptr {
				inputSamples := unsafe.Slice((*int16)(input), int(frameCount)*c.Capture.Channels)
				outputSamples := unsafe.Slice((*int32)(output), int(frameCount)*c.Playback.Channels)

				gotSamples := callback(inputSamples, int(frameCount), c.Playback.Channels, c.Capture.Channels)

				for i := range int(frameCount) {
					for j := range c.Playback.Channels {
						outputSamples[i*c.Playback.Channels+j] = gotSamples[i][j]
					}
				}

				return 0
			}, nil
		case FormatInt32:
			callbackPtr := (*DuplexCallback[int32, int32])(unsafe.Pointer(c.dataCallback.Load()))
			if callbackPtr == nil {
				return nil, ErrNilCallback
			}

			callback := *callbackPtr

			return func(_ *ma.Device, output, input unsafe.Pointer, frameCount uint32) uintptr {
				inputSamples := unsafe.Slice((*int32)(input), int(frameCount)*c.Capture.Channels)
				outputSamples := unsafe.Slice((*int32)(output), int(frameCount)*c.Playback.Channels)

				gotSamples := callback(inputSamples, int(frameCount), c.Playback.Channels, c.Capture.Channels)

				for i := range int(frameCount) {
					for j := range c.Playback.Channels {
						outputSamples[i*c.Playback.Channels+j] = gotSamples[i][j]
					}
				}

				return 0
			}, nil
		case FormatFloat32:
			callbackPtr := (*DuplexCallback[int32, float32])(unsafe.Pointer(c.dataCallback.Load()))
			if callbackPtr == nil {
				return nil, ErrNilCallback
			}

			callback := *callbackPtr

			return func(_ *ma.Device, output, input unsafe.Pointer, frameCount uint32) uintptr {
				inputSamples := unsafe.Slice((*float32)(input), int(frameCount)*c.Capture.Channels)
				outputSamples := unsafe.Slice((*int32)(output), int(frameCount)*c.Playback.Channels)

				gotSamples := callback(inputSamples, int(frameCount), c.Playback.Channels, c.Capture.Channels)

				for i := range int(frameCount) {
					for j := range c.Playback.Channels {
						outputSamples[i*c.Playback.Channels+j] = gotSamples[i][j]
					}
				}

				return 0
			}, nil
		default:
			return nil, ErrFormatNotSupported
		}
	case FormatFloat32:
		switch c.Capture.Format {
		case FormatUint8:
			callbackPtr := (*DuplexCallback[float32, uint8])(unsafe.Pointer(c.dataCallback.Load()))
			if callbackPtr == nil {
				return nil, ErrNilCallback
			}

			callback := *callbackPtr

			return func(_ *ma.Device, output, input unsafe.Pointer, frameCount uint32) uintptr {
				inputSamples := unsafe.Slice((*uint8)(input), int(frameCount)*c.Capture.Channels)
				outputSamples := unsafe.Slice((*float32)(output), int(frameCount)*c.Playback.Channels)

				gotSamples := callback(inputSamples, int(frameCount), c.Playback.Channels, c.Capture.Channels)

				for i := range int(frameCount) {
					for j := range c.Playback.Channels {
						outputSamples[i*c.Playback.Channels+j] = gotSamples[i][j]
					}
				}

				return 0
			}, nil
		case FormatInt16:
			callbackPtr := (*DuplexCallback[float32, int16])(unsafe.Pointer(c.dataCallback.Load()))
			if callbackPtr == nil {
				return nil, ErrNilCallback
			}

			callback := *callbackPtr

			return func(_ *ma.Device, output, input unsafe.Pointer, frameCount uint32) uintptr {
				inputSamples := unsafe.Slice((*int16)(input), int(frameCount)*c.Capture.Channels)
				outputSamples := unsafe.Slice((*float32)(output), int(frameCount)*c.Playback.Channels)

				gotSamples := callback(inputSamples, int(frameCount), c.Playback.Channels, c.Capture.Channels)

				for i := range int(frameCount) {
					for j := range c.Playback.Channels {
						outputSamples[i*c.Playback.Channels+j] = gotSamples[i][j]
					}
				}

				return 0
			}, nil
		case FormatInt32:
			callbackPtr := (*DuplexCallback[float32, int32])(unsafe.Pointer(c.dataCallback.Load()))
			if callbackPtr == nil {
				return nil, ErrNilCallback
			}

			callback := *callbackPtr

			return func(_ *ma.Device, output, input unsafe.Pointer, frameCount uint32) uintptr {
				inputSamples := unsafe.Slice((*int32)(input), int(frameCount)*c.Capture.Channels)
				outputSamples := unsafe.Slice((*float32)(output), int(frameCount)*c.Playback.Channels)

				gotSamples := callback(inputSamples, int(frameCount), c.Playback.Channels, c.Capture.Channels)

				for i := range int(frameCount) {
					for j := range c.Playback.Channels {
						outputSamples[i*c.Playback.Channels+j] = gotSamples[i][j]
					}
				}

				return 0
			}, nil
		case FormatFloat32:
			callbackPtr := (*DuplexCallback[float32, float32])(unsafe.Pointer(c.dataCallback.Load()))
			if callbackPtr == nil {
				return nil, ErrNilCallback
			}

			callback := *callbackPtr

			return func(_ *ma.Device, output, input unsafe.Pointer, frameCount uint32) uintptr {
				inputSamples := unsafe.Slice((*float32)(input), int(frameCount)*c.Capture.Channels)
				outputSamples := unsafe.Slice((*float32)(output), int(frameCount)*c.Playback.Channels)

				gotSamples := callback(inputSamples, int(frameCount), c.Playback.Channels, c.Capture.Channels)

				for i := range int(frameCount) {
					for j := range c.Playback.Channels {
						outputSamples[i*c.Playback.Channels+j] = gotSamples[i][j]
					}
				}

				return 0
			}, nil
		default:
			return nil, ErrFormatNotSupported
		}
	default:
		return nil, ErrFormatNotSupported
	}
}
