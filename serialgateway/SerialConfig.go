package serialgateway

// SerialConfig is onfiguration of serial port
type SerialConfig struct {
	PortNames       string
	BaudRate        uint
	DataBits        uint
	StopBits        uint
	MinimumReadSize uint
}
