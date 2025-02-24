package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"sync"

	"github.com/samborkent/miniaudio"
)

const bufferSize = 514

type sampleBuffer struct {
	samples [][]float32
}

var samplePool = sync.Pool{
	New: func() any {
		return &sampleBuffer{
			samples: make([][]float32, bufferSize),
		}
	},
}

var transferChannel = make(chan float32, bufferSize)

func captureCallback(inputSamples []float32) {
	for _, sample := range inputSamples {
		transferChannel <- sample
	}
}

func playbackBallback(frameCount, channelCount int) [][]float32 {
	buffer := samplePool.Get()

	samples, _ := buffer.(*sampleBuffer)

	samples.samples = make([][]float32, frameCount)

	for i := range frameCount {
		samples.samples[i] = make([]float32, channelCount)

		sample := <-transferChannel

		for c := range channelCount {
			samples.samples[i][c] = sample
		}
	}

	samplePool.Put(samples)

	return samples.samples
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	if err := miniaudio.Init(); err != nil {
		slog.ErrorContext(ctx, "initializing miniaudio: "+err.Error())
		return
	}

	device, err := miniaudio.NewDevice(miniaudio.DeviceConfig{
		DeviceType:       miniaudio.DeviceTypeDuplex,
		PlaybackCallback: playbackBallback,
		CaptureCallback:  captureCallback,
	})
	if err != nil {
		slog.ErrorContext(ctx, "creating new device: "+err.Error())
		return
	}

	if err := device.Init(); err != nil {
		slog.ErrorContext(ctx, "initializing device: "+err.Error())
		return
	}
	defer device.Uninit()

	slog.Info("initialized device", slog.Any("device", device))

	if err := device.Start(); err != nil {
		slog.ErrorContext(ctx, "starting device: "+err.Error())
		return
	}

	<-ctx.Done()

	slog.Info("bye")
}
