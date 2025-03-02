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

func (f Format) toMA() ma.Format {
	switch f {
	case FormatUint8:
		return ma.FormatU8
	case FormatInt16:
		return ma.FormatS16
	case FormatInt32:
		return ma.FormatS32
	case FormatFloat32:
		return ma.FormatF32
	default:
		return ma.FormatUnknown
	}
}

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
