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

	device := miniaudio.NewDevice(miniaudio.DeviceConfig{
		DeviceType: miniaudio.DeviceTypeCapture,
	})

	if err := device.Init(); err != nil {
		panic(err)
	}
	defer device.Uninit()

	slog.Info("initialized device", slog.Any("device", device))

	if err := device.Start(); err != nil {
		panic(err)
	}

	<-ctx.Done()

	slog.Info("bye")
}
