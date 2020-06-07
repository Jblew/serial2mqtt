package serialgateway

// Start starts serial gateway and connects the serial port
// and starts listening sending events to the channel
func Start(config Config, clock Clock, outputChan chan Event) SerialGateway {
	gateway := SerialGateway{}
	gateway.config = config
	gateway.clock = clock
	gateway.outputChan = outputChan

	go gateway.connectReadLoop()

	return gateway
}
