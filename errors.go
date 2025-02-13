package miniaudio

import (
	"errors"

	"github.com/samborkent/miniaudio/internal/ma"
)

var (
	ErrGeneric                    = errors.New("Error")
	ErrInvalidArgs                = errors.New("InvalidArgs")
	ErrInvalidOperation           = errors.New("InvalidOperation")
	ErrOutOfMemory                = errors.New("OutOfMemory")
	ErrOutOfRange                 = errors.New("OutOfRange")
	ErrAccessDenied               = errors.New("AccessDenied")
	ErrDoesNotExist               = errors.New("DoesNotExist")
	ErrAlreadyExists              = errors.New("AlreadyExists")
	ErrTooManyOpenFiles           = errors.New("TooManyOpenFiles")
	ErrInvalidFile                = errors.New("InvalidFile")
	ErrTooBig                     = errors.New("TooBig")
	ErrPathTooLong                = errors.New("PathTooLong")
	ErrNameTooLong                = errors.New("NameTooLong")
	ErrNotDirectory               = errors.New("NotDirectory")
	ErrIsDirectory                = errors.New("IsDirectory")
	ErrDirectoryNotEmpty          = errors.New("DirectoryNotEmpty")
	ErrAtEnd                      = errors.New("AtEnd")
	ErrNoSpace                    = errors.New("NoSpace")
	ErrBusy                       = errors.New("Busy")
	ErrIOError                    = errors.New("IOError")
	ErrInterrupt                  = errors.New("Interrupt")
	ErrUnavailable                = errors.New("Unavailable")
	ErrAlreadyInUse               = errors.New("AlreadyInUse")
	ErrBadAddress                 = errors.New("BadAddress")
	ErrBadSeek                    = errors.New("BadSeek")
	ErrBadPipe                    = errors.New("BadPipe")
	ErrDeadlock                   = errors.New("Deadlock")
	ErrTooManyLinks               = errors.New("TooManyLinks")
	ErrNotImplemented             = errors.New("NotImplemented")
	ErrNoMessage                  = errors.New("NoMessage")
	ErrBadMessage                 = errors.New("BadMessage")
	ErrNoDataAvailable            = errors.New("NoDataAvailable")
	ErrInvalidData                = errors.New("InvalidData")
	ErrTimeout                    = errors.New("Timeout")
	ErrNoNetwork                  = errors.New("NoNetwork")
	ErrNotUnique                  = errors.New("NotUnique")
	ErrNotSocket                  = errors.New("NotSocket")
	ErrNoAddress                  = errors.New("NoAddress")
	ErrBadProtocol                = errors.New("BadProtocol")
	ErrProtocolUnavailable        = errors.New("ProtocolUnavailable")
	ErrProtocolNotSupported       = errors.New("ProtocolNotSupported")
	ErrProtocolFamilyNotSupported = errors.New("ProtocolFamilyNotSupported")
	ErrAddressFamilyNotSupported  = errors.New("AddressFamilyNotSupported")
	ErrSocketNotSupported         = errors.New("SocketNotSupported")
	ErrConnectionReset            = errors.New("ConnectionReset")
	ErrAlreadyConnected           = errors.New("AlreadyConnected")
	ErrNotConnected               = errors.New("NotConnected")
	ErrConnectionRefused          = errors.New("ConnectionRefused")
	ErrNoHost                     = errors.New("NoHost")
	ErrInProgress                 = errors.New("InProgress")
	ErrCancelled                  = errors.New("Cancelled")
	ErrMemoryAlreadyMapped        = errors.New("MemoryAlreadyMapped")

	// General non-standard errors.
	ErrCRCMismatch = errors.New("CRCMismatch")

	// General miniaudio-specific errors.
	ErrFormatNotSupported     = errors.New("FormatNotSupported")
	ErrDeviceTypeNotSupported = errors.New("DeviceTypeNotSupported")
	ErrShareModeNotSupported  = errors.New("ShareModeNotSupported")
	ErrNoBackend              = errors.New("NoBackend")
	ErrNoDevice               = errors.New("NoDevice")
	ErrAPINotFound            = errors.New("APINotFound")
	ErrInvalidDeviceConfig    = errors.New("InvalidDeviceConfig")
	ErrLoop                   = errors.New("Loop")
	ErrBackendNotEnabled      = errors.New("BackendNotEnabled")

	// State errors.
	ErrDeviceNotInitialized     = errors.New("DeviceNotInitialized")
	ErrDeviceAlreadyInitialized = errors.New("DeviceAlreadyInitialized")
	ErrDeviceNotStarted         = errors.New("DeviceNotStarted")
	ErrDeviceNotStopped         = errors.New("DeviceNotStopped")

	// Operation errors.
	ErrFailedToInitBackend        = errors.New("FailedToInitBackend")
	ErrFailedToOpenBackendDevice  = errors.New("FailedToOpenBackendDevice")
	ErrFailedToStartBackendDevice = errors.New("FailedToStartBackendDevice")
	ErrFailedToStopBackendDevice  = errors.New("FailedToStopBackendDevice")
)

// TODO: finish
func convertResult(result ma.Result) error {
	switch result {
	case ma.Error:
		return ErrGeneric
	case ma.Success:
		fallthrough
	default:
		return nil
	}
}
