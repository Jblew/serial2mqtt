package serialgateway

import (
	"io"
)

// SerialGateway gateway class
type SerialGateway struct {
	clock             Clock
	config            Config
	outputChan        chan Event
	publishChan       chan []byte
	currentConnection io.ReadWriteCloser
}
