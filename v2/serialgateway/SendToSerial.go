package serialgateway

// SendToSerial publishes bytes to serial
func (gateway *SerialGateway) SendToSerial(payload []byte) error {
	gateway.publisher.AddToQueue(payload)
	return nil
}
