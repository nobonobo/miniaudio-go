package miniaudio

import (
	"unsafe"

	"github.com/samborkent/miniaudio/internal/ma"
)

type dataCallback func(device *ma.Device, output, input unsafe.Pointer, frameCount uint32) uintptr

func (c DeviceConfig[T]) playbackCallback(channelCount int) dataCallback {
	return func(_ *ma.Device, output, _ unsafe.Pointer, frameCount uint32) uintptr {
		outputSamples := unsafe.Slice((*T)(output), int(frameCount)*channelCount)

		gotSamples := c.PlaybackCallback(int(frameCount), channelCount)

		for i := range int(frameCount) {
			for c := range channelCount {
				outputSamples[i*channelCount+c] = gotSamples[i][c]
			}
		}

		return 0
	}
}

func (c DeviceConfig[T]) captureCallback(channelCount int) dataCallback {
	return func(_ *ma.Device, _, input unsafe.Pointer, frameCount uint32) uintptr {
		inputSamples := unsafe.Slice((*T)(input), int(frameCount)*channelCount)

		c.CaptureCallback(inputSamples, int(frameCount), channelCount)

		return 0
	}
}

func (c DeviceConfig[T]) duplexCallback(channelCount int) dataCallback {
	return func(_ *ma.Device, output, input unsafe.Pointer, frameCount uint32) uintptr {
		inputSamples := unsafe.Slice((*T)(input), int(frameCount)*channelCount)
		outputSamples := unsafe.Slice((*T)(output), int(frameCount)*channelCount)

		go c.CaptureCallback(inputSamples, int(frameCount), channelCount)
		gotSamples := c.PlaybackCallback(int(frameCount), channelCount)

		for i := range int(frameCount) {
			for c := range channelCount {
				outputSamples[i*channelCount+c] = gotSamples[i][c]
			}
		}

		return 0
	}
}
