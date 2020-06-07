package serialgateway

// Config is onfiguration of serial port
type Config struct {
	PortNames       string
	BaudRate        uint
	DataBits        uint
	StopBits        uint
	MinimumReadSize uint
}
