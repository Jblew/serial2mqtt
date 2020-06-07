package serialgateway

// Config is configuration of serial port
type Config struct {
	PortNames       string
	BaudRate        uint
	DataBits        uint
	StopBits        uint
	MinimumReadSize uint
}
