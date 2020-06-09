package serialgateway

import (
	"io"

	"github.com/Jblew/serial2mqtt/v2/serialgateway/publisher"
)

// SerialGateway gateway class
type SerialGateway struct {
	clock             Clock
	config            Config
	outputChan        chan Event
	currentConnection io.ReadWriteCloser
	publisher         *publisher.Publisher
}
