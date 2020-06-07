package serialgateway

import "fmt"

func (gateway *serialGateway) SendToSerial(payload []byte) error {
	if gateway.currentConnection == nil {
		return fmt.Errorf("Cannot publish. Serial connection not yet initialized")
	}

	_, err := gateway.currentConnection.Write(payload)
	if err != nil {
		return err
	}
	return nil
}
