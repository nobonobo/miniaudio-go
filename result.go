package miniaudio

// ma_result
type Result int32

const (
	Success                    Result = 0
	Error                      Result = -1 // A generic error.
	InvalidArgs                Result = -2
	InvalidOperation           Result = -3
	OutOfMemory                Result = -4
	OutOfRange                 Result = -5
	AccessDenied               Result = -6
	DoesNotExist               Result = -7
	AlreadyExists              Result = -8
	TooManyOpenFiles           Result = -9
	InvalidFile                Result = -10
	TooBig                     Result = -11
	PathTooLong                Result = -12
	NameTooLong                Result = -13
	NotDirectory               Result = -14
	IsDirectory                Result = -15
	DirectoryNotEmpty          Result = -16
	AtEnd                      Result = -17
	NoSpace                    Result = -18
	Busy                       Result = -19
	IOError                    Result = -20
	Interrupt                  Result = -21
	Unavailable                Result = -22
	AlreadyInUse               Result = -23
	BadAddress                 Result = -24
	BadSeek                    Result = -25
	BadPipe                    Result = -26
	Deadlock                   Result = -27
	TooManyLinks               Result = -28
	NotImplemented             Result = -29
	NoMessage                  Result = -30
	BadMessage                 Result = -31
	NoDataAvailable            Result = -32
	InvalidData                Result = -33
	Timeout                    Result = -34
	NoNetwork                  Result = -35
	NotUnique                  Result = -36
	NotSocket                  Result = -37
	NoAddress                  Result = -38
	BadProtocol                Result = -39
	ProtocolUnavailable        Result = -40
	ProtocolNotSupported       Result = -41
	ProtocolFamilyNotSupported Result = -42
	AddressFamilyNotSupported  Result = -43
	SocketNotSupported         Result = -44
	ConnectionReset            Result = -45
	AlreadyConnected           Result = -46
	NotConnected               Result = -47
	ConnectionRefused          Result = -48
	NoHost                     Result = -49
	InProgress                 Result = -50
	Cancelled                  Result = -51
	MemoryAlreadyMapped        Result = -52

	// General non-standard errors.
	CRCMismatch Result = -100

	// General miniaudio-specific errors.
	FormatNotSupported     Result = -200
	DeviceTypeNotSupported Result = -201
	ShareModeNotSupported  Result = -202
	NoBackend              Result = -203
	NoDevice               Result = -204
	APINotFound            Result = -205
	InvalidDeviceConfig    Result = -206
	Loop                   Result = -207
	BackendNotEnabled      Result = -208

	// State errors.
	DeviceNotInitialized     Result = -300
	DeviceAlreadyInitialized Result = -301
	DeviceNotStarted         Result = -302
	DeviceNotStopped         Result = -303

	// Operation errors.
	FailedToInitBackend        Result = -400
	FailedToOpenBackendDevice  Result = -401
	FailedToStartBackendDevice Result = -402
	FailedToStopBackendDevice  Result = -403
)
