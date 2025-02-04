package miniaudio

type maDeviceID[CustomType int | [256]byte | uintptr] struct {
	Wasapi      [64]uint16 // Equivalent to ma_wchar_win32 (wide character string).
	Dsound      [16]byte   // DirectSound uses a GUID for identification.
	Winmm       uint32     // WinMM expects a Win32 UINT_PTR (in practice a UINT).
	Alsa        [256]byte  // ALSA uses a name string for identification.
	Pulse       [256]byte  // PulseAudio uses a name string for identification.
	Jack        int        // JACK always uses default devices.
	CoreAudio   [256]byte  // Core Audio uses a string for identification.
	Sndio       [256]byte  // "snd/0", etc.
	Audio4      [256]byte  // "/dev/audio", etc.
	Oss         [64]byte   // "dev/dsp0", etc.
	Aaudio      int32      // AAudio uses a 32-bit integer for identification.
	Opensl      uint32     // OpenSL|ES uses a 32-bit unsigned integer for identification.
	WebAudio    [32]byte   // Web Audio uses default devices for now (potential GUID).
	Custom      CustomType // Custom backend options.
	NullBackend int        // Null backend uses an integer for device IDs.
}
