package serialgateway

import (
	"log"
)

// SendToSerial publishes bytes to serial
func (gateway *SerialGateway) SendToSerial(payload []byte) error {
	log.Printf("<><><> Current connection accessed %v", gateway.currentConnection)

	/*if gateway.currentConnection == nil {
		return fmt.Errorf("Cannot publish. Serial connection not yet initialized")
	}*/

	_, err := gateway.currentConnection.Write(payload)
	if err != nil {
		return err
	}
	return nil
}
