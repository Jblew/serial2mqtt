package serialgateway

import "serialsender/ntpclock"

// Run starts serial gateway and connects the serial port
// and starts listening sending events to the channel
func Run(config Config, clock ntpclock.NtpClock, outputChan chan Event) {
	gateway := serialGateway{}
	gateway.config = config
	gateway.clock = clock
	gateway.outputChan = outputChan

	gateway.connectReadLoop()
}
