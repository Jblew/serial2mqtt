package mqttqueue

// MQTTPublish is an interface that allows publishing MQTT messages
type MQTTPublish func(topic string, payload []byte) error
