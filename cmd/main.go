package main

import (
	"log/slog"

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
	// config.DataCallback = 0 // TODO: set callback function

	slog.Info("got device config", slog.Any("config", config))

	//     config.dataCallback      = data_callback;   // This function will be called when miniaudio needs more data.
	//     config.pUserData         = pMyCustomData;   // Can be accessed from the device object (device.pUserData).

	//     ma_device device;
	//     if (ma_device_init(NULL, &config, &device) != MA_SUCCESS) {
	//         return -1;  // Failed to initialize the device.
	//     }
}
