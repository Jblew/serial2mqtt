package serial2mqtt

import (
	"serialsender/serial2mqtt/mqttqueue"
	"serialsender/serial2mqtt/serialgateway"
)

// Serial2MQTT is main class for serial2mqtt lib
type Serial2MQTT struct {
	config  Config
	queue   mqttqueue.MQTTQueue
	gateway serialgateway.SerialGateway
}