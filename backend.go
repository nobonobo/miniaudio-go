package miniaudio

import "github.com/samborkent/miniaudio/internal/ma"

type Backend string

const (
	BackendWASAPI     Backend = "wasapi"
	BackendPulseAudio Backend = "pulseaudio"
	BackendALSA       Backend = "alsa"
	BackendJack       Backend = "jack"
)

func backendFromMA(backend ma.Backend) (Backend, error) {
	switch backend {
	case ma.BackendWASAPI:
		return BackendWASAPI, nil
	case ma.BackendPulseAudio:
		return BackendPulseAudio, nil
	case ma.BackendALSA:
		return BackendALSA, nil
	case ma.BackendJack:
		return BackendJack, nil
	default:
		return "", ErrBackendNotSupported
	}
}
