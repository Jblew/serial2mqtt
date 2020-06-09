package serialgateway

import (
	"fmt"
	"log"
)

// SendToSerial publishes bytes to serial
func (gateway *SerialGateway) SendToSerial(payload []byte) error {
	log.Printf("<><><> Current connection accessed %v", gateway.currentConnection)

	if gateway.publisher == nil {
		return fmt.Errorf("Cannot publish. Serial publisher not yet initialized")
	}

	err := gateway.publisher.Publish(payload)
	if err != nil {
		return err
	}
	return nil
}
