package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"sync"

	"github.com/samborkent/miniaudio"
)

const bufferSize = 2048

type sampleBuffer struct {
	samples [][]int16
}

var samplePool = sync.Pool{
	New: func() any {
		return &sampleBuffer{
			samples: make([][]int16, 0, bufferSize),
		}
	},
}

var transferChannel = make(chan int16, bufferSize)

func captureCallback(inputSamples []int16, frameCount, channelCount int) {
	for _, sample := range inputSamples {
		transferChannel <- sample
	}
}

func playbackBallback(frameCount, channelCount int) [][]int16 {
	buffer := samplePool.Get()

	samples, _ := buffer.(*sampleBuffer)
	clear(samples.samples)
	samples.samples = make([][]int16, frameCount)

	for i := range frameCount {
		samples.samples[i] = make([]int16, channelCount)

		sample := <-transferChannel

		for c := range channelCount {
			samples.samples[i][c] = sample
		}
	}

	samplePool.Put(samples)

	return samples.samples
}

func duplexCallback(inputSamples []int16, frameCount, playbackChannels, captureChannels int) [][]int16 {
	go captureCallback(inputSamples, frameCount, captureChannels)
	return playbackBallback(frameCount, playbackChannels)
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	if err := miniaudio.Init(); err != nil {
		slog.ErrorContext(ctx, "initializing miniaudio: "+err.Error())
		return
	}

	context := miniaudio.NewContext()

	if err := context.Init(); err != nil {
		slog.ErrorContext(ctx, "initializing miniaudio context: "+err.Error())
		return
	}
	defer context.Uninit()

	slog.DebugContext(ctx, "initialized context", slog.Any("context", context))

	playbackInfo, err := context.GetDefaultPlayback()
	if err != nil {
		slog.ErrorContext(ctx, "getting playback device: "+err.Error())
		return
	}

	playbackInfo, err = context.GetDeviceInfo(miniaudio.DeviceTypePlayback, playbackInfo.ID)
	if err != nil {
		slog.ErrorContext(ctx, "getting playback device info: "+err.Error())
		return
	}

	slog.InfoContext(ctx, "Device info", slog.Any("playback", playbackInfo))

	captureInfo, err := context.GetDefaultCapture()
	if err != nil {
		slog.ErrorContext(ctx, "getting capture device: "+err.Error())
		return
	}

	captureInfo, err = context.GetDeviceInfo(miniaudio.DeviceTypeCapture, captureInfo.ID)
	if err != nil {
		slog.ErrorContext(ctx, "getting capture device info: "+err.Error())
		return
	}

	slog.InfoContext(ctx, "Device info", slog.Any("capture", captureInfo))

	if len(playbackInfo.DataFormats) != 1 {
		slog.ErrorContext(ctx, "more than one native playback format")
		return
	}

	if len(captureInfo.DataFormats) != 1 {
		slog.ErrorContext(ctx, "more than one native capture format")
		return
	}

	deviceConfig := miniaudio.DeviceConfig{
		DeviceType: miniaudio.DeviceTypeDuplex,
		Playback: miniaudio.FormatConfig{
			Format:   playbackInfo.DataFormats[0].Format,
			Channels: playbackInfo.DataFormats[0].Channels,
		},
		Capture: miniaudio.FormatConfig{
			Format:   captureInfo.DataFormats[0].Format,
			Channels: captureInfo.DataFormats[0].Channels,
		},
	}

	miniaudio.SetDuplexCallback(&deviceConfig, duplexCallback)

	device, err := miniaudio.NewDevice(&deviceConfig)
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

	slog.Info("Device has started...")

	<-ctx.Done()

	slog.Info("Device is stopping...")

	if err := device.Stop(); err != nil {
		slog.ErrorContext(ctx, "stopping device: "+err.Error())
		return
	}

	slog.Info("bye")
}
