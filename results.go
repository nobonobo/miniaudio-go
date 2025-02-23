package miniaudio

import (
	"errors"

	"github.com/samborkent/miniaudio/internal/ma"
)

var (
	ErrResultGeneric                    = errors.New("Error")
	ErrResultInvalidArgs                = errors.New("InvalidArgs")
	ErrResultInvalidOperation           = errors.New("InvalidOperation")
	ErrResultOutOfMemory                = errors.New("OutOfMemory")
	ErrResultOutOfRange                 = errors.New("OutOfRange")
	ErrResultAccessDenied               = errors.New("AccessDenied")
	ErrResultDoesNotExist               = errors.New("DoesNotExist")
	ErrResultAlreadyExists              = errors.New("AlreadyExists")
	ErrResultTooManyOpenFiles           = errors.New("TooManyOpenFiles")
	ErrResultInvalidFile                = errors.New("InvalidFile")
	ErrResultTooBig                     = errors.New("TooBig")
	ErrResultPathTooLong                = errors.New("PathTooLong")
	ErrResultNameTooLong                = errors.New("NameTooLong")
	ErrResultNotDirectory               = errors.New("NotDirectory")
	ErrResultIsDirectory                = errors.New("IsDirectory")
	ErrResultDirectoryNotEmpty          = errors.New("DirectoryNotEmpty")
	ErrResultAtEnd                      = errors.New("AtEnd")
	ErrResultNoSpace                    = errors.New("NoSpace")
	ErrResultBusy                       = errors.New("Busy")
	ErrResultIOError                    = errors.New("IOError")
	ErrResultInterrupt                  = errors.New("Interrupt")
	ErrResultUnavailable                = errors.New("Unavailable")
	ErrResultAlreadyInUse               = errors.New("AlreadyInUse")
	ErrResultBadAddress                 = errors.New("BadAddress")
	ErrResultBadSeek                    = errors.New("BadSeek")
	ErrResultBadPipe                    = errors.New("BadPipe")
	ErrResultDeadlock                   = errors.New("Deadlock")
	ErrResultTooManyLinks               = errors.New("TooManyLinks")
	ErrResultNotImplemented             = errors.New("NotImplemented")
	ErrResultNoMessage                  = errors.New("NoMessage")
	ErrResultBadMessage                 = errors.New("BadMessage")
	ErrResultNoDataAvailable            = errors.New("NoDataAvailable")
	ErrResultInvalidData                = errors.New("InvalidData")
	ErrResultTimeout                    = errors.New("Timeout")
	ErrResultNoNetwork                  = errors.New("NoNetwork")
	ErrResultNotUnique                  = errors.New("NotUnique")
	ErrResultNotSocket                  = errors.New("NotSocket")
	ErrResultNoAddress                  = errors.New("NoAddress")
	ErrResultBadProtocol                = errors.New("BadProtocol")
	ErrResultProtocolUnavailable        = errors.New("ProtocolUnavailable")
	ErrResultProtocolNotSupported       = errors.New("ProtocolNotSupported")
	ErrResultProtocolFamilyNotSupported = errors.New("ProtocolFamilyNotSupported")
	ErrResultAddressFamilyNotSupported  = errors.New("AddressFamilyNotSupported")
	ErrResultSocketNotSupported         = errors.New("SocketNotSupported")
	ErrResultConnectionReset            = errors.New("ConnectionReset")
	ErrResultAlreadyConnected           = errors.New("AlreadyConnected")
	ErrResultNotConnected               = errors.New("NotConnected")
	ErrResultConnectionRefused          = errors.New("ConnectionRefused")
	ErrResultNoHost                     = errors.New("NoHost")
	ErrResultInProgress                 = errors.New("InProgress")
	ErrResultCancelled                  = errors.New("Cancelled")
	ErrResultMemoryAlreadyMapped        = errors.New("MemoryAlreadyMapped")

	// General non-standard errors.
	ErrResultCRCMismatch = errors.New("CRCMismatch")

	// General miniaudio-specific errors.
	ErrResultFormatNotSupported     = errors.New("FormatNotSupported")
	ErrResultDeviceTypeNotSupported = errors.New("DeviceTypeNotSupported")
	ErrResultShareModeNotSupported  = errors.New("ShareModeNotSupported")
	ErrResultNoBackend              = errors.New("NoBackend")
	ErrResultNoDevice               = errors.New("NoDevice")
	ErrResultAPINotFound            = errors.New("APINotFound")
	ErrResultInvalidDeviceConfig    = errors.New("InvalidDeviceConfig")
	ErrResultLoop                   = errors.New("Loop")
	ErrResultBackendNotEnabled      = errors.New("BackendNotEnabled")

	// State errors.
	ErrResultDeviceNotInitialized     = errors.New("DeviceNotInitialized")
	ErrResultDeviceAlreadyInitialized = errors.New("DeviceAlreadyInitialized")
	ErrResultDeviceNotStarted         = errors.New("DeviceNotStarted")
	ErrResultDeviceNotStopped         = errors.New("DeviceNotStopped")

	// Operation errors.
	ErrResultFailedToInitBackend        = errors.New("FailedToInitBackend")
	ErrResultFailedToOpenBackendDevice  = errors.New("FailedToOpenBackendDevice")
	ErrResultFailedToStartBackendDevice = errors.New("FailedToStartBackendDevice")
	ErrResultFailedToStopBackendDevice  = errors.New("FailedToStopBackendDevice")
)

func convertResult(result ma.Result) error {
	switch result {
	case ma.Error:
		return ErrResultGeneric
	case ma.InvalidArgs:
		return ErrResultInvalidArgs
	case ma.InvalidOperation:
		return ErrResultInvalidOperation
	case ma.OutOfMemory:
		return ErrResultOutOfMemory
	case ma.OutOfRange:
		return ErrResultOutOfRange
	case ma.AccessDenied:
		return ErrResultAccessDenied
	case ma.DoesNotExist:
		return ErrResultDoesNotExist
	case ma.AlreadyExists:
		return ErrResultAlreadyExists
	case ma.TooManyOpenFiles:
		return ErrResultTooManyOpenFiles
	case ma.InvalidFile:
		return ErrResultInvalidFile
	case ma.TooBig:
		return ErrResultTooBig
	case ma.PathTooLong:
		return ErrResultPathTooLong
	case ma.NameTooLong:
		return ErrResultNameTooLong
	case ma.NotDirectory:
		return ErrResultNotDirectory
	case ma.IsDirectory:
		return ErrResultIsDirectory
	case ma.DirectoryNotEmpty:
		return ErrResultDirectoryNotEmpty
	case ma.AtEnd:
		return ErrResultAtEnd
	case ma.NoSpace:
		return ErrResultNoSpace
	case ma.Busy:
		return ErrResultBusy
	case ma.IOError:
		return ErrResultIOError
	case ma.Interrupt:
		return ErrResultInterrupt
	case ma.Unavailable:
		return ErrResultUnavailable
	case ma.AlreadyInUse:
		return ErrResultAlreadyInUse
	case ma.BadAddress:
		return ErrResultBadAddress
	case ma.BadSeek:
		return ErrResultBadSeek
	case ma.BadPipe:
		return ErrResultBadPipe
	case ma.Deadlock:
		return ErrResultDeadlock
	case ma.TooManyLinks:
		return ErrResultTooManyLinks
	case ma.NotImplemented:
		return ErrResultNotImplemented
	case ma.NoMessage:
		return ErrResultNoMessage
	case ma.BadMessage:
		return ErrResultBadMessage
	case ma.NoDataAvailable:
		return ErrResultNoDataAvailable
	case ma.InvalidData:
		return ErrResultInvalidData
	case ma.Timeout:
		return ErrResultTimeout
	case ma.NoNetwork:
		return ErrResultNoNetwork
	case ma.NotUnique:
		return ErrResultNotUnique
	case ma.NotSocket:
		return ErrResultNotSocket
	case ma.NoAddress:
		return ErrResultNoAddress
	case ma.BadProtocol:
		return ErrResultBadProtocol
	case ma.ProtocolUnavailable:
		return ErrResultProtocolUnavailable
	case ma.ProtocolNotSupported:
		return ErrResultProtocolNotSupported
	case ma.ProtocolFamilyNotSupported:
		return ErrResultProtocolFamilyNotSupported
	case ma.AddressFamilyNotSupported:
		return ErrResultAddressFamilyNotSupported
	case ma.SocketNotSupported:
		return ErrResultSocketNotSupported
	case ma.ConnectionReset:
		return ErrResultConnectionReset
	case ma.AlreadyConnected:
		return ErrResultAlreadyConnected
	case ma.NotConnected:
		return ErrResultNotConnected
	case ma.ConnectionRefused:
		return ErrResultConnectionRefused
	case ma.NoHost:
		return ErrResultNoHost
	case ma.InProgress:
		return ErrResultInProgress
	case ma.Cancelled:
		return ErrResultCancelled
	case ma.MemoryAlreadyMapped:
		return ErrResultMemoryAlreadyMapped
	case ma.CRCMismatch:
		return ErrResultCRCMismatch
	case ma.FormatNotSupported:
		return ErrResultFormatNotSupported
	case ma.DeviceTypeNotSupported:
		return ErrResultDeviceTypeNotSupported
	case ma.ShareModeNotSupported:
		return ErrResultShareModeNotSupported
	case ma.NoBackend:
		return ErrResultNoBackend
	case ma.NoDevice:
		return ErrResultNoDevice
	case ma.APINotFound:
		return ErrResultAPINotFound
	case ma.InvalidDeviceConfig:
		return ErrResultInvalidDeviceConfig
	case ma.Loop:
		return ErrResultLoop
	case ma.BackendNotEnabled:
		return ErrResultBackendNotEnabled
	case ma.DeviceNotInitialized:
		return ErrResultDeviceNotInitialized
	case ma.DeviceAlreadyInitialized:
		return ErrResultDeviceAlreadyInitialized
	case ma.DeviceNotStarted:
		return ErrResultDeviceNotStarted
	case ma.DeviceNotStopped:
		return ErrResultDeviceNotStopped
	case ma.FailedToInitBackend:
		return ErrResultFailedToInitBackend
	case ma.FailedToOpenBackendDevice:
		return ErrResultFailedToOpenBackendDevice
	case ma.FailedToStartBackendDevice:
		return ErrResultFailedToStartBackendDevice
	case ma.FailedToStopBackendDevice:
		return ErrResultFailedToStopBackendDevice
	case ma.Success:
		fallthrough
	default:
		return nil
	}
}
