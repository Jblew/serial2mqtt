package serialgateway

import (
	"io"
	"serialsender/ntpclock"
)

// serialGateway gateway class
type serialGateway struct {
	clock             ntpclock.NtpClock
	config            Config
	outputChan        chan Event
	currentConnection io.ReadWriteCloser
}
