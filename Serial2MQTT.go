package serial2mqtt

import (
	"serial2mqtt/mqttqueue"
	"serial2mqtt/serialgateway"
)

// Serial2MQTT is main class for serial2mqtt lib
type Serial2MQTT struct {
	config  Config
	queue   mqttqueue.MQTTQueue
	gateway serialgateway.SerialGateway
}
