package main

import (
	"log/slog"
	"os"
	"runtime"

	"github.com/samborkent/miniaudio"
)

// void data_callback(ma_device* pDevice, void* pOutput, const void* pInput, ma_uint32 frameCount)
//     {
//         // In playback mode copy data to pOutput. In capture mode read data from pInput. In full-duplex mode, both
//         // pOutput and pInput will be valid and you can move data from pInput into pOutput. Never process more than
//         // frameCount frames.
//     }

func main() {
	config := miniaudio.DeviceConfigInit(miniaudio.DeviceTypePlayback)

	config.Playback.Format = miniaudio.FormatF32
	config.Playback.Channels = 2
	config.SampleRate = 48000
	// config.DataCallback = 0 // TODO: set callback function. This function will be called when miniaudio needs more data.

	slog.Info("got device config", slog.Any("config", config))

	var device *miniaudio.Device

	runtime.KeepAlive(device)

	result := miniaudio.DeviceInit(nil, config, device)
	if result != miniaudio.Success {
		os.Exit(int(result))
	}

	slog.Info("got device", slog.Any("device", device))

	result = miniaudio.DeviceStart(device)
	if result != miniaudio.Success {
		os.Exit(int(result))
	}

	// Do something here. Probably your program's main loop.

	miniaudio.DeviceUninit(device)

	os.Exit(0)
}
