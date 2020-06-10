package serial2mqtt

// SendToSerial publishes bytes to serial
func (serial2mqtt *Serial2MQTT) SendToSerial(payload []byte) error {
	serial2mqtt.publishChan <- payload
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
