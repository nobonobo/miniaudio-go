package miniaudio

type DeviceDescriptor struct {
	DeviceID                 *DeviceID
	ShareMode                ShareMode
	Format                   Format
	Channels                 uint32
	SampleRate               uint32
	ChannelMap               [MaxChannels]uint32 // Assuming channel map is an array of channel identifiers.
	PeriodSizeInFrames       uint32
	PeriodSizeInMilliseconds uint32
	PeriodCount              uint32
}
