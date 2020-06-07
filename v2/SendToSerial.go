package serial2mqtt

// SendToSerial publishes bytes to serial
func (serial2mqtt *Serial2MQTT) SendToSerial(payload []byte) error {
	err := serial2mqtt.gateway.SendToSerial(payload)
	if err != nil {
		return err
	}
	return nil
}

// SendToSerialStr publishes string to serial
func (serial2mqtt *Serial2MQTT) SendToSerialStr(payload string) error {
	err := serial2mqtt.SendToSerial([]byte(payload))
	if err != nil {
		return err
	}
	return nil
}
