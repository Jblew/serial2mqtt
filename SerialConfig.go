package serial2mqtt

// SerialConfig is configuration of serial port
type SerialConfig struct {
	PortNames       string
	BaudRate        uint
	DataBits        uint
	StopBits        uint
	MinimumReadSize uint
}
