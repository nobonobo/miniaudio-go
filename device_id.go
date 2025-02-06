package miniaudio

type DeviceID struct {
	WASAPI [64]uint64 // WASAPI uses a wchar_t string for identification
	DSound uint8      // DirectSound uses a GUID for identification
	WinMM  uint32     // WinMM expects a UINT_PTR for device identification
	ALSA   [256]uint8 // ALSA uses a name string for identification
	Pulse  [256]uint8 // PulseAudio uses a name string for identification
	Jack   int32      // JACK uses an integer for device ID
	// TODO: reanable when supporting additional platforms
	// CoreAudio [256]uint8 // Core Audio uses a string for identification
	// Sndio     [256]uint8 // "snd/0", etc.
	// Audio4    [256]uint8 // "/dev/audio", etc.
	// Oss       [64]uint8  // "dev/dsp0", etc.
	// AAudio    int32      // AAudio uses a 32-bit integer for identification
	// OpenSL    uint32     // OpenSL|ES uses a 32-bit unsigned integer for identification
	// WebAudio  [32]byte   // Web Audio always uses default devices for now
	// Custom    struct {
	// 	I int32
	// 	S [256]uint8
	// 	P unsafe.Pointer
	// } // The custom backend could be anything. Give them a few options.
	// NullBackend int32 // The null backend uses an integer for device ID
}
