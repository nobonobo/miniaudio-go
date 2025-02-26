package miniaudio

import "github.com/samborkent/miniaudio/internal/ma"

type Formats interface {
	uint8 | int16 | int32 | float32
}

type Format string

const (
	FormatUint8   Format = "uint8"
	FormatInt16   Format = "int16"
	FormatInt32   Format = "int32"
	FormatFloat32 Format = "float32"
)

func formatFromMA(format ma.Format) Format {
	switch format {
	case ma.FormatU8:
		return FormatUint8
	case ma.FormatS16:
		return FormatInt16
	case ma.FormatS32:
		return FormatInt32
	case ma.FormatF32:
		return FormatFloat32
	default:
		return ""
	}
}
