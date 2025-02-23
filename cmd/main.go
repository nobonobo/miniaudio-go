package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	"github.com/samborkent/miniaudio"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	if err := miniaudio.Init(); err != nil {
		slog.ErrorContext(ctx, "initializing miniaudio: "+err.Error())
		return
	}

	device, err := miniaudio.NewDevice(miniaudio.DeviceConfig{
		DeviceType: miniaudio.DeviceTypeDuplex,
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
