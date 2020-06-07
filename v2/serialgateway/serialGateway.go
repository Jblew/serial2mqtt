package serialgateway

import (
	"io"
)

// SerialGateway gateway class
type SerialGateway struct {
	clock             Clock
	config            Config
	outputChan        chan Event
	currentConnection io.ReadWriteCloser
}
