package serial2mqtt

// MQTTClient is an interface that allows publishing MQTT messages
type MQTTClient interface {
	// Publish method should block until sending is done
	// and return error on failure
	Publish(topic string, payload []byte) error
}
