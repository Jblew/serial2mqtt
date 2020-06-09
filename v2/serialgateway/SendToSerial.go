package serialgateway

import (
	"log"
)

// SendToSerial publishes bytes to serial
func (gateway *SerialGateway) SendToSerial(payload []byte) error {
	log.Printf("<><><> Current connection accessed %v", gateway.currentConnection)

	err := gateway.publisher.Publish(payload)
	if err != nil {
		return err
	}
	return nil
}
