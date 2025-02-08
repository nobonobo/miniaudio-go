package miniaudio

// void data_callback(ma_device* pDevice, void* pOutput, const void* pInput, ma_uint32 frameCount)
//     {
//         // In playback mode copy data to pOutput. In capture mode read data from pInput. In full-duplex mode, both
//         // pOutput and pInput will be valid and you can move data from pInput into pOutput. Never process more than
//         // frameCount frames.
//     }

func DeviceConfigInit(deviceType DeviceType) DeviceConfig {
	var deviceConfig DeviceConfig
	maDeviceConfigInit(&deviceConfig, deviceType)
	return deviceConfig
}
