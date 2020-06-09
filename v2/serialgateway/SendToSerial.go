package serialgateway

import (
	"fmt"
	"log"
)

// SendToSerial publishes bytes to serial
func (gateway *SerialGateway) SendToSerial(payload []byte) error {
	conn := *gateway.currentConnection
	log.Printf("<><><> Current connection accessed %v", conn)

	if conn == nil {
		return fmt.Errorf("Cannot publish. Serial connection not yet initialized")
	}

	_, err := conn.Write(payload)
	if err != nil {
		return err
	}
	return nil
}
