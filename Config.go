package serial2mqtt

import "serialsender/serial2mqtt/mqttqueue"

// Config for serial2mqtt gateway
type Config struct {
	MQTT mqttqueue.MQTTClient
}
